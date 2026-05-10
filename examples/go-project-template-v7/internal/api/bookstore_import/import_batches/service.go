package import_batches

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkSourceName string, pkBatchNo int) (*Dto, error)
	DeleteByID(ctx context.Context, pkSourceName string, pkBatchNo int) error
	FindByID(ctx context.Context, pkSourceName string, pkBatchNo int) (*Dto, error)
	FindAll(ctx context.Context) ([]Dto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Dto, pageable.Page, error)
}

type svc struct {
	repo Repository
}

var _ Service = &svc{}

func NewService(_ context.Context, repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) Save(ctx context.Context, input *CreateDto) (*Dto, error) {
	save, err := s.repo.Save(ctx, fromCreateDtoToEntity(input))
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(save)
	return &result, nil
}

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkSourceName string, pkBatchNo int) (*Dto, error) {
	updated, err := s.repo.UpdateByID(ctx, input, pkSourceName, pkBatchNo)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updated)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkSourceName string, pkBatchNo int) error {
	return s.repo.DeleteByID(ctx, pkSourceName, pkBatchNo)
}

func (s *svc) FindByID(ctx context.Context, pkSourceName string, pkBatchNo int) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkSourceName, pkBatchNo)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(entityByID)
	return &result, nil
}

func (s *svc) FindAll(ctx context.Context) ([]Dto, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return fromEntitiesToDtos(entities), nil
}

func (s *svc) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Dto, pageable.Page, error) {
	entities, page, err := s.repo.FindAllPageable(ctx, pq)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return fromEntitiesToDtos(entities), page, nil
}

// mappers

func fromCreateDtoToEntity(input *CreateDto) *ImportBatches {
	return &ImportBatches{
		SourceName: input.SourceName,
		BatchNo:    input.BatchNo,
		StartedAt:  input.StartedAt,
		FinishedAt: input.FinishedAt,
		FileName:   input.FileName,
		RowCount:   input.RowCount,
		Metadata:   input.Metadata,
	}
}

func fromEntitiesToDtos(inputEntities []ImportBatches) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *ImportBatches) Dto {
	return Dto{
		SourceName: inputEntity.SourceName,
		BatchNo:    inputEntity.BatchNo,
		StartedAt:  inputEntity.StartedAt,
		FinishedAt: inputEntity.FinishedAt,
		FileName:   inputEntity.FileName,
		RowCount:   inputEntity.RowCount,
		Metadata:   inputEntity.Metadata,
	}
}
