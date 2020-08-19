package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/smeruelo/glow/graph/generated"
	"github.com/smeruelo/glow/graph/model"
)

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	p1 := model.Project{
		Name:     "p1",
		Goal:     10,
		Achieved: 0,
	}
	p2 := model.Project{
		Name:     "p2",
		Goal:     80,
		Achieved: 0,
	}
	return []*model.Project{&p1, &p2}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
