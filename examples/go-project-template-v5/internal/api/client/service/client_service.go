package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/client/dto"
)

type ClientService interface {
	Save(ctx context.Context, input *dto.ClientCreateDto) (*dto.ClientDto, error)
	UpdateByID(ctx context.Context, pkRecordID int, input *dto.ClientUpdateDto) (*dto.ClientDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.ClientDto, error)
	FindAll(ctx context.Context) ([]dto.ClientDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ClientDto, pageable.Page, error)
}
