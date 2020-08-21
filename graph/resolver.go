package graph

import (
	"github.com/insta-code1/go-redis-replica/pkg/blueprints"
	"github.com/insta-code1/go-redis-replica/service"
)

// Resolver is.
type Resolver struct {
	Svc        *service.Service
	Blueprints *blueprints.Blueprints
}
