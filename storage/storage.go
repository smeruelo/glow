//go:generate mockery --name Store

package storage

import "github.com/smeruelo/glow/graph/model"

// Store defines the interface for projects storage
type Store interface {
	Create(project model.Project) error
	Get(id string) (model.Project, error)
	GetAll() ([]model.Project, error)
	Delete(id string) error
	UpdateAchieved(id string, time int) (model.Project, error)
}
