package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/currencies/dto"
)

type CurrenciesService interface {
	Save(ctx context.Context, input *dto.CurrenciesCreateDto) (*dto.CurrenciesDto, error)
	UpdateByID(ctx context.Context, input *dto.CurrenciesUpdateDto, pkRecordID int) (*dto.CurrenciesDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.CurrenciesDto, error)
	FindAll(ctx context.Context) ([]dto.CurrenciesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CurrenciesDto, pageable.Page, error)
}
