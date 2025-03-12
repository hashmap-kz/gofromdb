package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/customer_order_items/dto"
	dbModel "go-project-template-v5/internal/api/customer_order_items/entity/postgres"

	"go-project-template-v5/internal/api/customer_order_items/repository"
	"go-project-template-v5/internal/api/customer_order_items/service"
)

type customerOrderItemsService struct {
	customerOrderItemsRepository repository.CustomerOrderItemsRepository
}

var _ service.CustomerOrderItemsService = &customerOrderItemsService{}

func NewCustomerOrderItemsService(_ context.Context, customerOrderItemsRepository repository.CustomerOrderItemsRepository) service.CustomerOrderItemsService {
	return &customerOrderItemsService{
		customerOrderItemsRepository: customerOrderItemsRepository,
	}
}

func (s *customerOrderItemsService) Save(ctx context.Context, input *dto.CustomerOrderItemsCreateDto) (*dto.CustomerOrderItemsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.customerOrderItemsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customerOrderItemsService) UpdateByID(ctx context.Context, input *dto.CustomerOrderItemsUpdateDto, pkRecordID int) (*dto.CustomerOrderItemsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.customerOrderItemsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customerOrderItemsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.customerOrderItemsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *customerOrderItemsService) FindByID(ctx context.Context, pkRecordID int) (*dto.CustomerOrderItemsDto, error) {
	entityByID, err := s.customerOrderItemsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customerOrderItemsService) FindAll(ctx context.Context) ([]dto.CustomerOrderItemsDto, error) {
	entities, err := s.customerOrderItemsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *customerOrderItemsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CustomerOrderItemsDto, pageable.Page, error) {
	entities, page, err := s.customerOrderItemsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CustomerOrderItemsCreateDto) (*dbModel.CustomerOrderItems, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CustomerOrderItemsCreateDto->CustomerOrderItems: input dto cannot be nil")
	}
	return &dbModel.CustomerOrderItems{
		CustomerOrderID: input.CustomerOrderID,
		ProductID:       input.ProductID,
		Quantity:        input.Quantity,
		Price:           input.Price,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CustomerOrderItemsUpdateDto) (*dbModel.CustomerOrderItems, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CustomerOrderItemsUpdateDto->CustomerOrderItems: input dto cannot be nil")
	}
	return &dbModel.CustomerOrderItems{
		CustomerOrderID: input.CustomerOrderID,
		ProductID:       input.ProductID,
		Quantity:        input.Quantity,
		Price:           input.Price,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.CustomerOrderItems) ([]dto.CustomerOrderItemsDto, error) {
	outputDtos := make([]CustomerOrderItemsDto, 0, len(inputEntities))
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.CustomerOrderItems) (dto.CustomerOrderItemsDto, error) {
	if inputEntity == nil {
		return dto.CustomerOrderItemsDto{}, fmt.Errorf("unexpected nil input for mapping between CustomerOrderItems->CustomerOrderItemsDto")
	}
	return dto.CustomerOrderItemsDto{
		RecordID:        inputEntity.RecordID,
		CustomerOrderID: inputEntity.CustomerOrderID,
		ProductID:       inputEntity.ProductID,
		Quantity:        inputEntity.Quantity,
		Price:           inputEntity.Price,
		CreatedAt:       inputEntity.CreatedAt,
		UpdatedAt:       inputEntity.UpdatedAt,
		GUID:            inputEntity.GUID,
	}, nil
}
