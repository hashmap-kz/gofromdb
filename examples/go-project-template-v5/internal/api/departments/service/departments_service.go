package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/departments/dto"
)

type DepartmentsService interface {
	Save(ctx context.Context, input *dto.DepartmentsCreateDto) (*dto.DepartmentsDto, error)
	UpdateByID(ctx context.Context, input *dto.DepartmentsUpdateDto, pkRecordID int) (*dto.DepartmentsDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.DepartmentsDto, error)
	FindAll(ctx context.Context) ([]dto.DepartmentsDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.DepartmentsDto, pageable.Page, error)
}
