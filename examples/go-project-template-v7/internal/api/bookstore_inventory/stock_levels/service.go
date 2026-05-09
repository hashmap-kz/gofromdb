package stock_levels

import (
	"context"
	"fmt"
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkWarehouseCode string, pkBookID int64) (*Dto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.repo.UpdateByID(ctx, entityToUpdate, pkWarehouseCode, pkBookID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *svc) DeleteByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) error {
	return s.repo.DeleteByID(ctx, pkWarehouseCode, pkBookID)
}

func (s *svc) FindByID(ctx context.Context, pkWarehouseCode string, pkBookID int64) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkWarehouseCode, pkBookID)
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

func fromCreateDtoToEntity(input *CreateDto) (*StockLevels, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CreateDto->StockLevels: input dto cannot be nil")
	}
	return &StockLevels{
		WarehouseCode:    input.WarehouseCode,
		BookID:           input.BookID,
		AvailableQty:     input.AvailableQty,
		ReservedQty:      input.ReservedQty,
		ReorderThreshold: input.ReorderThreshold,
		LastCountedAt:    input.LastCountedAt,
	}, nil
}

func fromUpdateDtoToEntity(input *UpdateDto) (*StockLevels, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UpdateDto->StockLevels: input dto cannot be nil")
	}
	return &StockLevels{
		AvailableQty:     input.AvailableQty,
		ReservedQty:      input.ReservedQty,
		ReorderThreshold: input.ReorderThreshold,
		LastCountedAt:    input.LastCountedAt,
	}, nil
}

func fromEntitiesToDtos(inputEntities []StockLevels) ([]Dto, error) {
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

func fromEntityToDto(inputEntity *StockLevels) (Dto, error) {
	if inputEntity == nil {
		return Dto{}, fmt.Errorf("unexpected nil input for mapping between StockLevels->Dto")
	}
	return Dto{
		WarehouseCode:    inputEntity.WarehouseCode,
		BookID:           inputEntity.BookID,
		AvailableQty:     inputEntity.AvailableQty,
		ReservedQty:      inputEntity.ReservedQty,
		ReorderThreshold: inputEntity.ReorderThreshold,
		LastCountedAt:    inputEntity.LastCountedAt,
	}, nil
}
