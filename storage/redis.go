package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/smeruelo/glow/graph/model"
)

type redisStore struct {
	conn redis.Conn
}

// NewRedisStore creates a Store that implements the interface for a Redis storage
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

func (s redisStore) Create(p model.Project) error {
	pJSON, err := json.Marshal(p)
	if err != nil {
		log.Printf("Unable to marshal project: %s", err)
		return err
	}
	n, err := redis.Int64(s.conn.Do("HSETNX", "projects", p.ID, pJSON))
	if err != nil {
		log.Printf("Database error: %s", err)
		return err
	}
	if n != 1 {
		log.Printf("Project %s already exists", p.ID)
		return fmt.Errorf("Project %s already exists", p.ID)
	}
	return nil
}

func (s redisStore) Get(id string) (model.Project, error) {
	var p model.Project
	pJSON, err := redis.Bytes(s.conn.Do("HGET", "projects", id))
	if err != nil {
		log.Printf("Database error: %s", err)
		return p, err
	}
	if err := json.Unmarshal(pJSON, &p); err != nil {
		log.Printf("Unable to unmarshal project: %s", err)
		return p, err
	}
	return p, nil
}

func (s redisStore) GetAll() ([]model.Project, error) {
	psJSON, err := redis.StringMap(s.conn.Do("HGETALL", "projects"))
	if err != nil {
		log.Printf("Database error: %s", err)
		return nil, err
	}

	ps := make([]model.Project, len(psJSON))
	i := 0
	for _, pJSON := range psJSON {
		var p model.Project
		if err := json.Unmarshal([]byte(pJSON), &p); err != nil {
			log.Printf("Unable to unmarshal project: %s", err)
			return ps, err
		}
		ps[i] = p
		i++
	}
	return ps, nil
}

func (s redisStore) Delete(id string) error {
	n, err := redis.Int64(s.conn.Do("HDEL", "projects", id))
	if err != nil {
		log.Printf("Database error: %s", err)
		return err
	}
	if n < 1 {
		log.Printf("Project %s does not exist", id)
		return fmt.Errorf("Project %s does not exist", id)
	}
	return nil
}

// ToDo: Solve possible race conditions
func (s redisStore) UpdateAchieved(id string, time int) (model.Project, error) {
	var p model.Project

	pJSON, err := redis.Bytes(s.conn.Do("HGET", "projects", id))
	if err != nil {
		log.Printf("Database error: %s", err)
		return p, err
	}
	if err := json.Unmarshal(pJSON, &p); err != nil {
		log.Printf("Unable to unmarshal project: %s", err)
		return p, err
	}

	p.Achieved += time
	pJSON, err = json.Marshal(p)
	if err != nil {
		log.Printf("Unable to marshal project: %s", err)
		return p, err
	}

	_, err = redis.Int64(s.conn.Do("HSET", "projects", id, pJSON))
	if err != nil {
		log.Printf("Database error: %s", err)
		return p, err
	}

	return p, nil
}
