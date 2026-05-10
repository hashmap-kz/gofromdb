package categories

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkRecordID int) (*Dto, error)
	DeleteByID(ctx context.Context, pkRecordID int) error
	FindByID(ctx context.Context, pkRecordID int) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkRecordID int) (*Dto, error) {
	updated, err := s.repo.UpdateByID(ctx, input, pkRecordID)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updated)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.repo.DeleteByID(ctx, pkRecordID)
}

func (s *svc) FindByID(ctx context.Context, pkRecordID int) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkRecordID)
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

func fromCreateDtoToEntity(input *CreateDto) *Categories {
	return &Categories{
		Name:        input.Name,
		ParentID:    input.ParentID,
		ValidPeriod: input.ValidPeriod,
	}
}

func fromEntitiesToDtos(inputEntities []Categories) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *Categories) Dto {
	return Dto{
		RecordID:    inputEntity.RecordID,
		Name:        inputEntity.Name,
		ParentID:    inputEntity.ParentID,
		ValidPeriod: inputEntity.ValidPeriod,
		IsCurrent:   inputEntity.IsCurrent,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		GUID:        inputEntity.GUID,
	}
}
