package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/client/dto"
)

type ClientService interface {
	Save(ctx context.Context, input *dto.ClientCreateDto) (*dto.ClientDto, error)
	GetAll(ctx context.Context) ([]dto.ClientDto, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ClientDto, pageable.Page, error)
	Update(ctx context.Context, entityId int, input *dto.ClientUpdateDto) (*dto.ClientDto, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.ClientDto, error)
}
