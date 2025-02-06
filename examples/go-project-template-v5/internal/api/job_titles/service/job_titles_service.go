package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/job_titles/dto"
)

type JobTitlesService interface {
	Save(ctx context.Context, input *dto.JobTitlesCreateDto) (*dto.JobTitlesDto, error)
	UpdateByID(ctx context.Context, input *dto.JobTitlesUpdateDto, pkRecordID int) (*dto.JobTitlesDto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dto.JobTitlesDto, error)
	FindAll(ctx context.Context) ([]dto.JobTitlesDto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.JobTitlesDto, pageable.Page, error)
}
