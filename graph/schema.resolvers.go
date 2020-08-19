package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/smeruelo/glow/graph/generated"
	"github.com/smeruelo/glow/graph/model"
)

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	projectsJSON, err := redis.StringMap(r.db.Do("HGETALL", "projects"))
	if err != nil {
		log.Printf("Database error: %s", err)
		return nil, err
	}

	projects := make([]*model.Project, len(projectsJSON))
	i := 0
	for id, pJSON := range projectsJSON {
		var p model.Project
		if err := json.Unmarshal([]byte(pJSON), &p); err != nil {
			log.Printf("Unable to unmarshal project: %s", err)
			return nil, err
		}
		p.ID = id
		projects[i] = &p
		i++
	}
	return projects, nil
}

func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
