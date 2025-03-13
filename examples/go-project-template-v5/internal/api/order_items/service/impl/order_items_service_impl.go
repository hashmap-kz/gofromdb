package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/order_items/dto"
	dbModel "go-project-template-v5/internal/api/order_items/entity/postgres"

	"go-project-template-v5/internal/api/order_items/repository"
	"go-project-template-v5/internal/api/order_items/service"
)

type orderItemsService struct {
	orderItemsRepository repository.OrderItemsRepository
}

var _ service.OrderItemsService = &orderItemsService{}

func NewOrderItemsService(_ context.Context, orderItemsRepository repository.OrderItemsRepository) service.OrderItemsService {
	return &orderItemsService{
		orderItemsRepository: orderItemsRepository,
	}
}

func (s *orderItemsService) Save(ctx context.Context, input *dto.OrderItemsCreateDto) (*dto.OrderItemsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.orderItemsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *orderItemsService) UpdateByID(ctx context.Context, input *dto.OrderItemsUpdateDto, pkRecordID int) (*dto.OrderItemsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.orderItemsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *orderItemsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.orderItemsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *orderItemsService) FindByID(ctx context.Context, pkRecordID int) (*dto.OrderItemsDto, error) {
	entityByID, err := s.orderItemsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *orderItemsService) FindAll(ctx context.Context) ([]dto.OrderItemsDto, error) {
	entities, err := s.orderItemsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *orderItemsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.OrderItemsDto, pageable.Page, error) {
	entities, page, err := s.orderItemsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.OrderItemsCreateDto) (*dbModel.OrderItems, error) {
	if input == nil {
		return nil, fmt.Errorf("convert OrderItemsCreateDto->OrderItems: input dto cannot be nil")
	}
	return &dbModel.OrderItems{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.OrderItemsUpdateDto) (*dbModel.OrderItems, error) {
	if input == nil {
		return nil, fmt.Errorf("convert OrderItemsUpdateDto->OrderItems: input dto cannot be nil")
	}
	return &dbModel.OrderItems{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.OrderItems) ([]dto.OrderItemsDto, error) {
	outputDtos := make([]dto.OrderItemsDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.OrderItems) (dto.OrderItemsDto, error) {
	if inputEntity == nil {
		return dto.OrderItemsDto{}, fmt.Errorf("unexpected nil input for mapping between OrderItems->OrderItemsDto")
	}
	return dto.OrderItemsDto{
		RecordID:  inputEntity.RecordID,
		OrderID:   inputEntity.OrderID,
		ProductID: inputEntity.ProductID,
		Quantity:  inputEntity.Quantity,
		Price:     inputEntity.Price,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		GUID:      inputEntity.GUID,
	}, nil
}
