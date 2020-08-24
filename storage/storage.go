//go:generate mockery --name Store

package storage

import "github.com/smeruelo/glow/graph/model"

// Store defines the interface for projects storage
type Store interface {
	CreateProject(p model.Project) error
	GetProject(pID string) (model.Project, error)
	GetUserProjects(uID string) ([]model.Project, error)
	DeleteProject(pID, uID string) error
	UpdateProject(pID string, np model.NewProject) (model.Project, error)
}
