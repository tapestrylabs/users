package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/tapestrylabs/users/api/config"
	"github.com/tapestrylabs/users/api/ent"
	"go.uber.org/zap"
)

type Resolver struct {
	client *ent.Client
	logger *zap.Logger
	config *config.ServerConfig
}

func NewSchema(client *ent.Client, logger *zap.Logger, config *config.ServerConfig) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			client,
			logger,
			config,
		},
	})
}
