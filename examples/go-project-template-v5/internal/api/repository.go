package api

import (
	"context"

	clientRepoInterface "go-project-template-v5/internal/api/client/repository"
	clientRepoImpl "go-project-template-v5/internal/api/client/repository/impl"

	"go-project-template-v5/pkg/storage/postgres"
)

// Init

type Repositories struct {
	// TODO: other repository interfaces here
	ClientRepository clientRepoInterface.ClientRepository
}

func NewRepositories(ctx context.Context, db *postgres.Postgres) *Repositories {
	return &Repositories{
		// TODO: other repository impls here
		ClientRepository: clientRepoImpl.NewClientRepository(ctx, db),
	}
}
