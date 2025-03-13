package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/orders/dto"
	dbModel "go-project-template-v5/internal/api/orders/entity/postgres"

	"go-project-template-v5/internal/api/orders/repository"
	"go-project-template-v5/internal/api/orders/service"
)

type ordersService struct {
	ordersRepository repository.OrdersRepository
}

var _ service.OrdersService = &ordersService{}

func NewOrdersService(_ context.Context, ordersRepository repository.OrdersRepository) service.OrdersService {
	return &ordersService{
		ordersRepository: ordersRepository,
	}
}

func (s *ordersService) Save(ctx context.Context, input *dto.OrdersCreateDto) (*dto.OrdersDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.ordersRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *ordersService) UpdateByID(ctx context.Context, input *dto.OrdersUpdateDto, pkRecordID int) (*dto.OrdersDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.ordersRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *ordersService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.ordersRepository.DeleteByID(ctx, pkRecordID)
}

func (s *ordersService) FindByID(ctx context.Context, pkRecordID int) (*dto.OrdersDto, error) {
	entityByID, err := s.ordersRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *ordersService) FindAll(ctx context.Context) ([]dto.OrdersDto, error) {
	entities, err := s.ordersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *ordersService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.OrdersDto, pageable.Page, error) {
	entities, page, err := s.ordersRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.OrdersCreateDto) (*dbModel.Orders, error) {
	if input == nil {
		return nil, fmt.Errorf("convert OrdersCreateDto->Orders: input dto cannot be nil")
	}
	return &dbModel.Orders{
		UserID:      input.UserID,
		Description: input.Description,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.OrdersUpdateDto) (*dbModel.Orders, error) {
	if input == nil {
		return nil, fmt.Errorf("convert OrdersUpdateDto->Orders: input dto cannot be nil")
	}
	return &dbModel.Orders{
		UserID:      input.UserID,
		Description: input.Description,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Orders) ([]dto.OrdersDto, error) {
	outputDtos := make([]dto.OrdersDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Orders) (dto.OrdersDto, error) {
	if inputEntity == nil {
		return dto.OrdersDto{}, fmt.Errorf("unexpected nil input for mapping between Orders->OrdersDto")
	}
	return dto.OrdersDto{
		RecordID:    inputEntity.RecordID,
		UserID:      inputEntity.UserID,
		Description: inputEntity.Description,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		GUID:        inputEntity.GUID,
	}, nil
}
