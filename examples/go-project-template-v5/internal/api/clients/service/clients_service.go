package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/clients/dto"
)

type ClientsService interface {
	Save(ctx context.Context, input *dto.ClientsCreateDto) (*dto.ClientsDto, error)
	UpdateByID(ctx context.Context, input *dto.ClientsUpdateDto, pkRecordID int) (*dto.ClientsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.ClientsDto, error)
	FindAll(ctx context.Context) ([]dto.ClientsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ClientsDto, pageable.Page, error)
}
