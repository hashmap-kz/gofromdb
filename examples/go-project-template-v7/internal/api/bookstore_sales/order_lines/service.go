package order_lines

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkOrderID int64, pkLineNo int) (*Dto, error)
	DeleteByID(ctx context.Context, pkOrderID int64, pkLineNo int) error
	FindByID(ctx context.Context, pkOrderID int64, pkLineNo int) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkOrderID int64, pkLineNo int) (*Dto, error) {
	updatedResult, err := s.repo.UpdateByID(ctx, fromUpdateDtoToEntity(input), pkOrderID, pkLineNo)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updatedResult)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkOrderID int64, pkLineNo int) error {
	return s.repo.DeleteByID(ctx, pkOrderID, pkLineNo)
}

func (s *svc) FindByID(ctx context.Context, pkOrderID int64, pkLineNo int) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkOrderID, pkLineNo)
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

func fromCreateDtoToEntity(input *CreateDto) *OrderLines {
	return &OrderLines{
		OrderID:        input.OrderID,
		LineNo:         input.LineNo,
		BookID:         input.BookID,
		Quantity:       input.Quantity,
		UnitPrice:      input.UnitPrice,
		DiscountAmount: input.DiscountAmount,
		Note:           input.Note,
	}
}

func fromUpdateDtoToEntity(input *UpdateDto) *OrderLines {
	return &OrderLines{
		BookID:         input.BookID,
		Quantity:       input.Quantity,
		UnitPrice:      input.UnitPrice,
		DiscountAmount: input.DiscountAmount,
		Note:           input.Note,
	}
}

func fromEntitiesToDtos(inputEntities []OrderLines) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *OrderLines) Dto {
	return Dto{
		OrderID:        inputEntity.OrderID,
		LineNo:         inputEntity.LineNo,
		BookID:         inputEntity.BookID,
		Quantity:       inputEntity.Quantity,
		UnitPrice:      inputEntity.UnitPrice,
		DiscountAmount: inputEntity.DiscountAmount,
		Note:           inputEntity.Note,
	}
}
