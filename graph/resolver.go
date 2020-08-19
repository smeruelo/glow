package graph

import "github.com/gomodule/redigo/redis"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db redis.Conn
}

func NewResolver(db redis.Conn) *Resolver {
	return &Resolver{db: db}
}
