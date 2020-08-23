//go:generate mockery --name Store

package storage

import "github.com/smeruelo/glow/graph/model"

// Store defines the interface for projects storage
type Store interface {
	CreateProject(project model.Project) error
	GetProject(id string) (model.Project, error)
	GetUserProjects(userID string) ([]model.Project, error)
	DeleteProject(id, userID string) error
	UpdateProject(id string, np model.NewProject) (model.Project, error)
}
