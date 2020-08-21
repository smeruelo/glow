package graph

import (
	"github.com/smeruelo/glow/storage"
)

// Resolver contains the dependencies needed to build the schema resolvers
type Resolver struct {
	store storage.Store
}

// NewResolver receives a DB store and creates a Resolver with it.
func NewResolver(s storage.Store) *Resolver {
	return &Resolver{store: s}
}
