package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/customer_orders/dto"
	dbModel "go-project-template-v5/internal/api/customer_orders/entity/postgres"

	"go-project-template-v5/internal/api/customer_orders/repository"
	"go-project-template-v5/internal/api/customer_orders/service"
)

type customerOrdersService struct {
	customerOrdersRepository repository.CustomerOrdersRepository
}

var _ service.CustomerOrdersService = &customerOrdersService{}

func NewCustomerOrdersService(_ context.Context, customerOrdersRepository repository.CustomerOrdersRepository) service.CustomerOrdersService {
	return &customerOrdersService{
		customerOrdersRepository: customerOrdersRepository,
	}
}

func (s *customerOrdersService) Save(ctx context.Context, input *dto.CustomerOrdersCreateDto) (*dto.CustomerOrdersDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.customerOrdersRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customerOrdersService) UpdateByID(ctx context.Context, input *dto.CustomerOrdersUpdateDto, pkRecordID int) (*dto.CustomerOrdersDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.customerOrdersRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customerOrdersService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.customerOrdersRepository.DeleteByID(ctx, pkRecordID)
}

func (s *customerOrdersService) FindByID(ctx context.Context, pkRecordID int) (*dto.CustomerOrdersDto, error) {
	entityByID, err := s.customerOrdersRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customerOrdersService) FindAll(ctx context.Context) ([]dto.CustomerOrdersDto, error) {
	entities, err := s.customerOrdersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *customerOrdersService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CustomerOrdersDto, pageable.Page, error) {
	entities, page, err := s.customerOrdersRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CustomerOrdersCreateDto) (*dbModel.CustomerOrders, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CustomerOrdersCreateDto->CustomerOrders: input dto cannot be nil")
	}
	return &dbModel.CustomerOrders{
		ClientID:    input.ClientID,
		Description: input.Description,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CustomerOrdersUpdateDto) (*dbModel.CustomerOrders, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CustomerOrdersUpdateDto->CustomerOrders: input dto cannot be nil")
	}
	return &dbModel.CustomerOrders{
		ClientID:    input.ClientID,
		Description: input.Description,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.CustomerOrders) ([]dto.CustomerOrdersDto, error) {
	outputDtos := make([]CustomerOrdersDto, 0, len(inputEntities))
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.CustomerOrders) (dto.CustomerOrdersDto, error) {
	if inputEntity == nil {
		return dto.CustomerOrdersDto{}, fmt.Errorf("unexpected nil input for mapping between CustomerOrders->CustomerOrdersDto")
	}
	return dto.CustomerOrdersDto{
		RecordID:    inputEntity.RecordID,
		ClientID:    inputEntity.ClientID,
		Description: inputEntity.Description,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		GUID:        inputEntity.GUID,
	}, nil
}
