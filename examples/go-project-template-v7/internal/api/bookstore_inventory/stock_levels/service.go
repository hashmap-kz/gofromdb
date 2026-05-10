package stock_levels

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkWarehouseCode string, pkBookID int64) (*Dto, error)
	DeleteByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) error
	FindByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkWarehouseCode string, pkBookID int64) (*Dto, error) {
	updatedResult, err := s.repo.UpdateByID(ctx, fromUpdateDtoToEntity(input), pkWarehouseCode, pkBookID)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updatedResult)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) error {
	return s.repo.DeleteByID(ctx, pkWarehouseCode, pkBookID)
}

func (s *svc) FindByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkWarehouseCode, pkBookID)
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

func fromCreateDtoToEntity(input *CreateDto) *StockLevels {
	return &StockLevels{
		WarehouseCode:    input.WarehouseCode,
		BookID:           input.BookID,
		AvailableQty:     input.AvailableQty,
		ReservedQty:      input.ReservedQty,
		ReorderThreshold: input.ReorderThreshold,
		LastCountedAt:    input.LastCountedAt,
	}
}

func fromUpdateDtoToEntity(input *UpdateDto) *StockLevels {
	return &StockLevels{
		AvailableQty:     input.AvailableQty,
		ReservedQty:      input.ReservedQty,
		ReorderThreshold: input.ReorderThreshold,
		LastCountedAt:    input.LastCountedAt,
	}
}

func fromEntitiesToDtos(inputEntities []StockLevels) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *StockLevels) Dto {
	return Dto{
		WarehouseCode:    inputEntity.WarehouseCode,
		BookID:           inputEntity.BookID,
		AvailableQty:     inputEntity.AvailableQty,
		ReservedQty:      inputEntity.ReservedQty,
		ReorderThreshold: inputEntity.ReorderThreshold,
		LastCountedAt:    inputEntity.LastCountedAt,
	}
}
