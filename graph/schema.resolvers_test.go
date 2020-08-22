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
		Name: "Test",
		Goal: 100,
	}
	p := model.Project{
		ID:       "",
		Name:     "Test",
		Goal:     100,
		Achieved: 0,
	}
	expected := &p

	s.On("Create", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
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
		Name: "Test",
		Goal: 100,
	}

	s.On("Create", mock.Anything).Return(errors.New(""))

	_, err := r.CreateProject(ctx, np)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestProjectSuccess(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	id := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	p := model.Project{
		ID:       id,
		Name:     "Test",
		Goal:     100,
		Achieved: 20,
	}
	expected := &p

	s.On("Get", id).Return(p, nil)

	actual, err := r.Project(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestProjectFail(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	id := "3b054f50-9d3d-4114-bfc4-395f70a59d26"

	s.On("Get", id).Return(model.Project{}, errors.New(""))

	_, err := r.Project(ctx, id)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestProjectsSuccess(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	p1 := model.Project{
		ID:       "3b054f50-9d3d-4114-bfc4-000000000001",
		Name:     "Test1",
		Goal:     100,
		Achieved: 10,
	}
	p2 := model.Project{
		ID:       "3b054f50-9d3d-4114-bfc4-000000000002",
		Name:     "Test2",
		Goal:     200,
		Achieved: 20,
	}
	expected := []*model.Project{&p1, &p2}

	s.On("GetAll").Return([]model.Project{p1, p2}, nil)

	actual, err := r.Projects(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestProjectsFails(t *testing.T) {
	var s mocks.Store
	r := &queryResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	s.On("GetAll").Return([]model.Project{}, errors.New(""))

	_, err := r.Projects(ctx)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestDeleteProjectSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	id := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	expected := id

	s.On("Delete", id).Return(nil)

	actual, err := r.DeleteProject(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestDeleteProjectFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	id := "3b054f50-9d3d-4114-bfc4-395f70a59d26"

	s.On("Delete", id).Return(errors.New(""))

	_, err := r.DeleteProject(ctx, id)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestUpdateProjectAchievedSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	id := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	time := 60
	p := model.Project{
		ID:       id,
		Name:     "Test",
		Goal:     100,
		Achieved: 80,
	}
	expected := &p

	s.On("UpdateAchieved", id, time).Return(p, nil)

	actual, err := r.UpdateProjectAchieved(ctx, id, time)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestUpdateProjectAchievedFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	id := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	time := 60
	var p model.Project

	s.On("UpdateAchieved", id, time).Return(p, errors.New(""))

	_, err := r.UpdateProjectAchieved(ctx, id, time)

	assert.Error(t, err)
	s.AssertExpectations(t)
}
