package repository

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	dbModel "go-project-template-v5/internal/api/job_titles/entity/postgres"
)

type JobTitlesRepository interface {
	Save(ctx context.Context, inputEntity *dbModel.JobTitles) (*dbModel.JobTitles, error)
	UpdateByID(ctx context.Context, inputEntity *dbModel.JobTitles, pkRecordID int) (*dbModel.JobTitles, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*dbModel.JobTitles, error)
	FindAll(ctx context.Context) ([]dbModel.JobTitles, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dbModel.JobTitles, pageable.Page, error)
}
