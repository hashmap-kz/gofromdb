package authors

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkAuthorID string) (*Dto, error)
	DeleteByID(ctx context.Context, pkAuthorID string) error
	FindByID(ctx context.Context, pkAuthorID string) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkAuthorID string) (*Dto, error) {
	updatedResult, err := s.repo.UpdateByID(ctx, fromUpdateDtoToEntity(input), pkAuthorID)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updatedResult)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkAuthorID string) error {
	return s.repo.DeleteByID(ctx, pkAuthorID)
}

func (s *svc) FindByID(ctx context.Context, pkAuthorID string) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkAuthorID)
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

func fromCreateDtoToEntity(input *CreateDto) *Authors {
	return &Authors{
		AuthorID:    input.AuthorID,
		DisplayName: input.DisplayName,
		LegalName:   input.LegalName,
		Biography:   input.Biography,
		Metadata:    input.Metadata,
		Active:      input.Active,
		BornOn:      input.BornOn,
	}
}

func fromUpdateDtoToEntity(input *UpdateDto) *Authors {
	return &Authors{
		DisplayName: input.DisplayName,
		LegalName:   input.LegalName,
		Biography:   input.Biography,
		Metadata:    input.Metadata,
		Active:      input.Active,
		BornOn:      input.BornOn,
	}
}

func fromEntitiesToDtos(inputEntities []Authors) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *Authors) Dto {
	return Dto{
		AuthorID:    inputEntity.AuthorID,
		DisplayName: inputEntity.DisplayName,
		LegalName:   inputEntity.LegalName,
		Biography:   inputEntity.Biography,
		Metadata:    inputEntity.Metadata,
		Active:      inputEntity.Active,
		BornOn:      inputEntity.BornOn,
	}
}
