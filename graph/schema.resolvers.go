package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/smeruelo/glow/graph/generated"
	"github.com/smeruelo/glow/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	p := model.Project{
		ID:       uuid.New().String(),
		Name:     input.Name,
		Goal:     input.Goal,
		Achieved: 0,
	}
	return &p, r.store.Create(p)
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id string) (string, error) {
	return id, r.store.Delete(id)
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	all, err := r.store.GetAll()
	if err != nil {
		return nil, err
	}
	ps := make([]*model.Project, len(all))
	for i := range all {
		ps[i] = &all[i]
	}
	return ps, nil
}

func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	p, err := r.store.Get(id)
	return &p, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
