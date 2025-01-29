package api

import (
	"context"

	clientServiceInterface "go-project-template-v5/internal/api/<no value>/service"
	clientServiceImpl "go-project-template-v5/internal/api/<no value>/service/impl"
)

// Init

type Services struct {
	// TODO: other service interfaces here
	ClientService clientServiceInterface.ClientService
}

type Deps struct {
	// TODO: other deps here
	Repos *Repositories
}

func NewServices(ctx context.Context, deps Deps) *Services {
	return &Services{
		// TODO: other service impls here
		ClientService: clientServiceImpl.NewClientService(ctx, deps.Repos.ClientRepository),
	}
}
