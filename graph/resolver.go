package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "graphql-sample/graph/services"

type Resolver struct {
	Srv services.Services
}
