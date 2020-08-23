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
		UserID:   "0",
		Name:     input.Name,
		Category: input.Category,
	}
	return &p, r.store.CreateProject(p)
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id string, userID string) (string, error) {
	return id, r.store.DeleteProject(id, "0")
}

func (r *mutationResolver) UpdateProject(ctx context.Context, id string, input model.NewProject) (*model.Project, error) {
	p, err := r.store.UpdateProject(id, input)
	return &p, err
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	all, err := r.store.GetUserProjects("0")
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
	p, err := r.store.GetProject(id)
	return &p, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
