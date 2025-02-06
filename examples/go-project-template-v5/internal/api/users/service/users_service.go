package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/users/dto"
)

type UsersService interface {
	Save(ctx context.Context, input *dto.UsersCreateDto) (*dto.UsersDto, error)
	UpdateByID(ctx context.Context, input *dto.UsersUpdateDto, pkRecordID int) (*dto.UsersDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.UsersDto, error)
	FindAll(ctx context.Context) ([]dto.UsersDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.UsersDto, pageable.Page, error)
}
