package storage

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/smeruelo/glow/graph/model"
)

type redisStore struct {
	conn redis.Conn
}

// NewRedisStore creates a Store that implements the interface for a Redis storage
// This storage implementation does not take into account race conditions yet
// Therefore, IT DOES NOT SUPPORT CONCURRENCY
//
// DB schema:
//
// |------------------------------|------------|------------------------------------------------------|
// | Key name                     | Redis type | Fields                                               |
// |------------------------------|------------|------------------------------------------------------|
// | users                        | hash       | email, userID                                        |
// | user:<userID>                | hash       | name, email, pass                                    |
// | sessions:<userID>            | set        | token                                                |
// | session:<token>              | hash       | userID                                               |
// | projects:<userID>            | set        | projectID                                            |
// | project:<projectID>          | hash       | userID, name, category                               |
// | achievements:<projectID>     | set        | achievementID                                        |
// | achievement:<achievementID>  | hash       | userID, projectID, startDateTime, endDateTime        |
// | goals:<projectID>            | set        | goalID                                               |
// | goal:<goalID>                | hash       | userID, projectID, type, minutes, startDate, endDate |
// |------------------------------|------------|------------------------------------------------------|
//
func NewRedisStore(conn redis.Conn) Store {
	return redisStore{conn: conn}
}

// Redis keys and hashes' fields
const (
	sCategory  string = "category"
	sName      string = "name"
	sProject   string = "project"
	sProjectID string = "projectID"
	sProjects  string = "projects"
	sUserID    string = "userID"
)

func (s redisStore) errIfDoesntExist(key string) error {
	n, err := redis.Int64(s.conn.Do("EXISTS", key))
	if err != nil {
		log.Printf("Database error: %s", err)
		return err
	}
	if n != 1 {
		log.Printf("Key %s does not exist", key)
		return fmt.Errorf("Key %s does not exist", key)
	}
	return nil
}

func (s redisStore) CreateProject(p model.Project) error {
	// Check if project exists
	key := fmt.Sprintf("%s:%s", sProject, p.ID)
	n, err := redis.Int64(s.conn.Do("EXISTS", key))
	if err != nil {
		log.Printf("Database error: %s", err)
		return err
	}
	if n == 1 {
		log.Printf("Project %s already exists", p.ID)
		return fmt.Errorf("Project %s already exists", p.ID)
	}

	// Create project
	_, err = redis.Int64(s.conn.Do("HSET", key, sUserID, p.UserID, sName, p.Name, sCategory, p.Category))
	if err != nil {
		log.Printf("Database error: %s", err)
		return err
	}

	// Add it to user's projects
	key = fmt.Sprintf("%s:%s", sProjects, p.UserID)
	_, err = redis.Int64(s.conn.Do("SADD", key, p.ID))
	if err != nil {
		log.Printf("Database error: %s", err)
		return err
	}

	return nil
}

func (s redisStore) getProject(pID string) (model.Project, error) {
	var p model.Project
	key := fmt.Sprintf("%s:%s", sProject, pID)
	if err := s.errIfDoesntExist(key); err != nil {
		return p, err
	}

	fields, err := redis.StringMap(s.conn.Do("HGETALL", key))
	if err != nil {
		log.Printf("Database error: %s", err)
		return p, err
	}

	p.ID = pID
	p.UserID = fields[sUserID]
	p.Name = fields[sName]
	p.Category = fields[sCategory]

	return p, nil
}

func (s redisStore) GetProject(pID string) (model.Project, error) {
	return s.getProject(pID)
}

func (s redisStore) GetUserProjects(uID string) ([]model.Project, error) {
	key := fmt.Sprintf("%s:%s", sProjects, uID)
	projectIDs, err := redis.Strings(s.conn.Do("SMEMBERS", key))
	if err != nil {
		log.Printf("Database error: %s", err)
		return nil, err
	}

	ps := make([]model.Project, len(projectIDs))
	for i, pID := range projectIDs {
		p, err := s.getProject(pID)
		if err != nil {
			return ps, err
		}
		ps[i] = p
	}
	return ps, nil
}

func (s redisStore) DeleteProject(pID, uID string) error {
	key := fmt.Sprintf("%s:%s", sProject, pID)
	if err := s.errIfDoesntExist(key); err != nil {
		return err
	}

	if _, err := redis.Int64(s.conn.Do("DEL", key)); err != nil {
		log.Printf("Database error: %s", err)
		return err
	}

	key = fmt.Sprintf("%s:%s", sProjects, uID)
	if _, err := redis.Int64(s.conn.Do("SREM", key, pID)); err != nil {
		log.Printf("Database error: %s", err)
		return err
	}

	return nil
}

func (s redisStore) UpdateProject(pID string, np model.NewProject) (model.Project, error) {
	p, err := s.getProject(pID)
	if err != nil {
		return p, err
	}

	key := fmt.Sprintf("%s:%s", sProject, pID)
	_, err = redis.Int64(s.conn.Do("HSET", key, sName, np.Name, sCategory, np.Category))
	if err != nil {
		log.Printf("Database error: %s", err)
		return p, err
	}

	p.Name = np.Name
	p.Category = np.Category
	return p, nil
}
