package publishers

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkCode string) (*Dto, error)
	DeleteByID(ctx context.Context, pkCode string) error
	FindByID(ctx context.Context, pkCode string) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkCode string) (*Dto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.repo.UpdateByID(ctx, entityToUpdate, pkCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *svc) DeleteByID(ctx context.Context, pkCode string) error {
	return s.repo.DeleteByID(ctx, pkCode)
}

func (s *svc) FindByID(ctx context.Context, pkCode string) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
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

func fromCreateDtoToEntity(input *CreateDto) (*Publishers, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CreateDto->Publishers: input dto cannot be nil")
	}
	return &Publishers{
		Code:        input.Code,
		Name:        input.Name,
		CountryCode: input.CountryCode,
		Website:     input.Website,
		FoundedOn:   input.FoundedOn,
		Active:      input.Active,
	}, nil
}

func fromUpdateDtoToEntity(input *UpdateDto) (*Publishers, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UpdateDto->Publishers: input dto cannot be nil")
	}
	return &Publishers{
		Name:        input.Name,
		CountryCode: input.CountryCode,
		Website:     input.Website,
		FoundedOn:   input.FoundedOn,
		Active:      input.Active,
	}, nil
}

func fromEntitiesToDtos(inputEntities []Publishers) ([]Dto, error) {
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

func fromEntityToDto(inputEntity *Publishers) (Dto, error) {
	if inputEntity == nil {
		return Dto{}, fmt.Errorf("unexpected nil input for mapping between Publishers->Dto")
	}
	return Dto{
		Code:        inputEntity.Code,
		Name:        inputEntity.Name,
		CountryCode: inputEntity.CountryCode,
		Website:     inputEntity.Website,
		FoundedOn:   inputEntity.FoundedOn,
		Active:      inputEntity.Active,
	}, nil
}
