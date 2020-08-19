package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
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
	for _, pJSON := range projectsJSON {
		var p model.Project
		if err := json.Unmarshal([]byte(pJSON), &p); err != nil {
			log.Printf("Unable to unmarshal project: %s", err)
			return nil, err
		}
		projects[i] = &p
		i++
	}
	return projects, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
