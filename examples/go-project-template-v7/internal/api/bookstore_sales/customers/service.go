package customers

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkCustomerID string) (*Dto, error)
	DeleteByID(ctx context.Context, pkCustomerID string) error
	FindByID(ctx context.Context, pkCustomerID string) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkCustomerID string) (*Dto, error) {
	updatedResult, err := s.repo.UpdateByID(ctx, fromUpdateDtoToEntity(input), pkCustomerID)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updatedResult)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkCustomerID string) error {
	return s.repo.DeleteByID(ctx, pkCustomerID)
}

func (s *svc) FindByID(ctx context.Context, pkCustomerID string) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkCustomerID)
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

func fromCreateDtoToEntity(input *CreateDto) *Customers {
	return &Customers{
		CustomerID:     input.CustomerID,
		Email:          input.Email,
		FullName:       input.FullName,
		Phone:          input.Phone,
		MarketingOptIn: input.MarketingOptIn,
		RegisteredAt:   input.RegisteredAt,
	}
}

func fromUpdateDtoToEntity(input *UpdateDto) *Customers {
	return &Customers{
		Email:          input.Email,
		FullName:       input.FullName,
		Phone:          input.Phone,
		MarketingOptIn: input.MarketingOptIn,
		RegisteredAt:   input.RegisteredAt,
	}
}

func fromEntitiesToDtos(inputEntities []Customers) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *Customers) Dto {
	return Dto{
		CustomerID:     inputEntity.CustomerID,
		Email:          inputEntity.Email,
		FullName:       inputEntity.FullName,
		Phone:          inputEntity.Phone,
		MarketingOptIn: inputEntity.MarketingOptIn,
		RegisteredAt:   inputEntity.RegisteredAt,
	}
}
