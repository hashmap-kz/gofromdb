package import_errors

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
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
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.repo.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *svc) FindAll(ctx context.Context) ([]Dto, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *svc) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Dto, pageable.Page, error) {
	entities, page, err := s.repo.FindAllPageable(ctx, pq)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return toDtos, page, nil
}

// mappers

func fromCreateDtoToEntity(input *CreateDto) (*ImportErrors, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CreateDto->ImportErrors: input dto cannot be nil")
	}
	return &ImportErrors{
		SourceName: input.SourceName,
		BatchNo:    input.BatchNo,
		RowNo:      input.RowNo,
		ColumnName: input.ColumnName,
		Message:    input.Message,
		RawPayload: input.RawPayload,
	}, nil
}

func fromEntitiesToDtos(inputEntities []ImportErrors) ([]Dto, error) {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *ImportErrors) (Dto, error) {
	if inputEntity == nil {
		return Dto{}, fmt.Errorf("unexpected nil input for mapping between ImportErrors->Dto")
	}
	return Dto{
		SourceName: inputEntity.SourceName,
		BatchNo:    inputEntity.BatchNo,
		RowNo:      inputEntity.RowNo,
		ColumnName: inputEntity.ColumnName,
		Message:    inputEntity.Message,
		RawPayload: inputEntity.RawPayload,
		CreatedAt:  inputEntity.CreatedAt,
	}, nil
}
