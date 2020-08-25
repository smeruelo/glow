package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

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

func (r *mutationResolver) UpdateProject(ctx context.Context, id string, input model.NewProject) (*model.Project, error) {
	p, err := r.store.UpdateProject(id, input)
	return &p, err
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id string) (string, error) {
	return id, r.store.DeleteProject(id, "0")
}

func (r *mutationResolver) CreateAchievement(ctx context.Context, projectID string) (*model.Achievement, error) {
	a := model.Achievement{
		ID:        uuid.New().String(),
		UserID:    "0",
		ProjectID: projectID,
		Start:     time.Now().Unix(),
		End:       0,
	}
	return &a, r.store.CreateAchievement(a)
}

func (r *mutationResolver) UpdateAchievement(ctx context.Context, id string, input model.AchievementData) (*model.Achievement, error) {
	a, err := r.store.UpdateAchievement(id, input)
	return &a, err
}

func (r *mutationResolver) DeleteAchievement(ctx context.Context, id string, projectID string) (string, error) {
	return id, r.store.DeleteProject(id, pID)
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

func (r *queryResolver) Achievement(ctx context.Context, id string) (*model.Achievement, error) {
	a, err := r.store.GetAchievement(id)
	return &a, err
}

func (r *queryResolver) ProjectAchievements(ctx context.Context, projectID string) ([]*model.Achievement, error) {
	all, err := r.store.GetProjectAchievements(projectID)
	if err != nil {
		return nil, err
	}
	as := make([]*model.Achievement, len(all))
	for i := range all {
		as[i] = &all[i]
	}
	return as, nil
}

func (r *queryResolver) UserAchievements(ctx context.Context) ([]*model.Achievement, error) {
	uID := "0"
	projects, err := r.store.GetUserProjects(uID)
	if err != nil {
		return nil, err
	}

	as := []*model.Achievement{}
	for p := range projects {
		pAs, err := r.store.GetProjectAchievements(p.ID)
		if err != nil {
			return nil, err
		}
		for a := range pAs {
			append(as, &a)
		}
	}

	return as, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
