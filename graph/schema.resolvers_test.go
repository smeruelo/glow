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

func TestUpdateProjectSuccess(t *testing.T) {
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

func TestUpdateProjectFail(t *testing.T) {
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

func TestCreateAchievementSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	a := model.Achievement{
		ID:        "",
		UserID:    "0",
		ProjectID: pID,
		Start:     1598341158,
		End:       0,
	}
	expected := &a

	s.On("CreateAchievement", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		ach := args.Get(0).(model.Achievement)
		a.ID = ach.ID
	})

	actual, err := r.CreateAchievement(pID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestCreateAchievementFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"

	s.On("CreateAchievement", mock.Anything).Return(errors.New(""))

	err := r.CreateAchievement(pID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestAchievementSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	aID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	a := model.Achievement{
		ID:        aID,
		UserID:    "0",
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b7df83",
		Start:     1598341158,
		End:       1598342861,
	}
	expected := &a

	s.On("GetAchievement", aID).Return(a, nil)

	actual, err := r.Query().Achievement(ctx, aID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestAchievementFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	aID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"

	s.On("GetAchievement", aID).Return(a, nil)

	_, err := r.Query().Achievement(ctx, aID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestProjectAchievementsSuccess(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "b1265627-d9f2-4a0b-b60d-322273b7df83"
	a1 := model.Achievement{
		ID:        "3b054f50-9d3d-4114-bfc4-395f70a00001",
		UserID:    "0",
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b00001",
		Start:     1598341158,
		End:       1598342861,
	}
	a2 := model.Achievement{
		ID:        "3b054f50-9d3d-4114-bfc4-395f70a00002",
		UserID:    "0",
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b00002",
		Start:     1598342900,
		End:       1598346500,
	}
	expected := []*model.Achievement{&a1, &a2}

	s.On("GetProjectAchievements", pID).Return([]model.Achievement{&a1, &a2}, nil)

	actual, err := r.Query().ProjectAchievements(ctx, pID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestProjectAchievementsFail(t *testing.T) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	pID := "b1265627-d9f2-4a0b-b60d-322273b7df83"

	s.On("GetProjectAchievements", pID).Return([]model.Achievement{}, errors.New(""))

	_, err := r.Query().ProjectAchievements(ctx, pID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestUserAchievementsSucces(t *testingT) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	uID := "0"
	a1 := model.Achievement{
		ID:        "3b054f50-9d3d-4114-bfc4-395f70a00001",
		UserID:    uID,
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b00001",
		Start:     1598341158,
		End:       1598342861,
	}
	a2 := model.Achievement{
		ID:        "3b054f50-9d3d-4114-bfc4-395f70a00002",
		UserID:    uID,
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b00002",
		Start:     1598342900,
		End:       1598346500,
	}
	expected := []*model.Achievement{&a1, &a2}

	s.On("GetUserAchievements", uID).Return([]model.Achievement{&a1, &a2}, nil)

	actual, err := r.Query().UserAchievements(ctx, uID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestUserAchievementsFail(t *testingT) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	uID := "0"

	s.On("GetUserAchievements", uID).Return([]model.Achievement{}, errors.New(""))

	_, err := r.Query().UserAchievements(ctx, uID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestUpdateAchievementSuccess(t *testingT) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	aID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	a := model.Achievement{
		ID:        aID,
		UserID:    "0",
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b7df83",
		Start:     1598341158,
		End:       1598342861,
	}
	ad := model.AchievementData{
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b7df83",
		Start:     1598341158,
		End:       1598342861,
	}
	expected := &a

	s.On("UpdateAchievement", aID, ad).Return(a, nil)

	actual, err := r.Mutation().UpdateAchievement(ctx, aID, ad)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestUpdateAchievementFail(t *testingT) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	aID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	ad := model.AchievementData{
		ProjectID: "b1265627-d9f2-4a0b-b60d-322273b7df83",
		Start:     1598341158,
		End:       1598342861,
	}

	s.On("UpdateAchievement", aID, ad).Return(model.Achievement{}, errors.New(""))

	_, err := r.Mutation().UpdateAchievement(ctx, aID, ad)

	assert.Error(t, err)
	s.AssertExpectations(t)
}

func TestDeleteAchievementSuccess(t *testingT) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	aID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	pID := "b1265627-d9f2-4a0b-b60d-322273b7df83"
	expected := aID

	s.On("DeleteAchievement", aID, pID).Return(nil)

	actual, err := r.Mutation().DeleteAchievement(ctx, aID, pID)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	s.AssertExpectations(t)
}

func TestDeleteAchievementFail(t *testingT) {
	var s mocks.Store
	r := &mutationResolver{Resolver: NewResolver(&s)}
	ctx := context.Background()

	aID := "3b054f50-9d3d-4114-bfc4-395f70a59d26"
	pID := "b1265627-d9f2-4a0b-b60d-322273b7df83"

	s.On("DeleteAchievement", aID, pID).Return(errors.New(""))

	_, err := r.Mutation().DeleteAchievement(ctx, aID, pID)

	assert.Error(t, err)
	s.AssertExpectations(t)
}
