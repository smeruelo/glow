package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/smeruelo/glow/graph/model"
	"github.com/smeruelo/glow/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProjectSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	np := model.NewProject{
		Name:     "Test",
		Category: "Default",
	}
	p := model.Project{
		ID:       "",
		UserID:   "0",
		Name:     "Test",
		Category: "Default",
	}
	expected := &p

	s.On("CreateProject", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		project := args.Get(0).(model.Project)
		p.ID = project.ID
	})

	actual, err := r.CreateProject(ctx, np)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestCreateProjectFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	np := model.NewProject{
		Name:     "Test",
		Category: "Default",
	}

	s.On("CreateProject", mock.Anything).Return(errors.New(""))

	_, err := r.CreateProject(ctx, np)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestProjectSuccess(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	p := model.Project{
		ID:       pID,
		UserID:   "0",
		Name:     "Test",
		Category: "Default",
	}
	expected := &p

	s.On("GetProject", pID).Return(p, nil)

	actual, err := r.Project(ctx, pID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestProjectFail(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"

	s.On("GetProject", pID).Return(model.Project{}, errors.New(""))

	_, err := r.Project(ctx, pID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestProjectsSuccess(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	p1 := model.Project{
		ID:       "3b054f50-9d3d-4114-bfc4-000000000001",
		UserID:   "0",
		Name:     "Test 1",
		Category: "Default",
	}
	p2 := model.Project{
		ID:       "3b054f50-9d3d-4114-bfc4-000000000002",
		UserID:   "0",
		Name:     "Test 2",
		Category: "Programming",
	}
	uID := "0"
	expected := []*model.Project{&p1, &p2}

	s.On("GetUserProjects", uID).Return([]model.Project{p1, p2}, nil)

	actual, err := r.Projects(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestProjectsFail(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	uID := "0"

	s.On("GetUserProjects", uID).Return([]model.Project{}, errors.New(""))

	_, err := r.Projects(ctx)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestDeleteProjectSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	uID := "0"
	expected := pID

	s.On("DeleteProject", pID, uID).Return(nil)

	actual, err := r.DeleteProject(ctx, pID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestDeleteProjectFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	uID := "0"

	s.On("DeleteProject", pID, uID).Return(errors.New(""))

	_, err := r.DeleteProject(ctx, pID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestUpdateProjectAchievedSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	p := model.Project{
		ID:       pID,
		UserID:   "0",
		Name:     "Test",
		Category: "Reading",
	}
	np := model.NewProject{
		Name:     "Test",
		Category: "Reading",
	}
	expected := &p

	s.On("UpdateProject", pID, np).Return(p, nil)

	actual, err := r.UpdateProject(ctx, pID, np)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestUpdateProjectAchievedFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	np := model.NewProject{
		Name:     "Test",
		Category: "Reading",
	}
	var p model.Project

	s.On("UpdateProject", pID, np).Return(p, errors.New(""))

	_, err := r.UpdateProject(ctx, pID, np)

	assert.Error(t, err)
	s.AssertExpectations(t)
}
