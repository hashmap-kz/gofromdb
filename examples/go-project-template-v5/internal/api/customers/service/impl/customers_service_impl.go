package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/customers/dto"
	dbModel "go-project-template-v5/internal/api/customers/entity/postgres"

	"go-project-template-v5/internal/api/customers/repository"
	"go-project-template-v5/internal/api/customers/service"
)

type customersService struct {
	customersRepository repository.CustomersRepository
}

var _ service.CustomersService = &customersService{}

func NewCustomersService(_ context.Context, customersRepository repository.CustomersRepository) service.CustomersService {
	return &customersService{
		customersRepository: customersRepository,
	}
}

func (s *customersService) Save(ctx context.Context, input *dto.CustomersCreateDto) (*dto.CustomersDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.customersRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customersService) UpdateByID(ctx context.Context, input *dto.CustomersUpdateDto, pkRecordID int) (*dto.CustomersDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.customersRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customersService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.customersRepository.DeleteByID(ctx, pkRecordID)
}

func (s *customersService) FindByID(ctx context.Context, pkRecordID int) (*dto.CustomersDto, error) {
	entityById, err := s.customersRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *customersService) FindAll(ctx context.Context) ([]dto.CustomersDto, error) {
	entities, err := s.customersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *customersService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CustomersDto, pageable.Page, error) {
	entities, page, err := s.customersRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CustomersCreateDto) (*dbModel.Customers, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CustomersCreateDto->Customers: input dto cannot be nil")
	}
	return &dbModel.Customers{
		Email: input.Email,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CustomersUpdateDto) (*dbModel.Customers, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CustomersUpdateDto->Customers: input dto cannot be nil")
	}
	return &dbModel.Customers{
		Email: input.Email,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Customers) ([]dto.CustomersDto, error) {
	var outputDtos []dto.CustomersDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Customers) (dto.CustomersDto, error) {
	if inputEntity == nil {
		return dto.CustomersDto{}, fmt.Errorf("unexpected nil input for mapping between Customers->CustomersDto")
	}
	return dto.CustomersDto{
		RecordID:  inputEntity.RecordID,
		Email:     inputEntity.Email,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}
