package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchases/dto"
	dbModel "go-project-template-v5/internal/api/purchases/entity/postgres"

	"go-project-template-v5/internal/api/purchases/repository"
	"go-project-template-v5/internal/api/purchases/service"
)

type purchasesService struct {
	purchasesRepository repository.PurchasesRepository
}

var _ service.PurchasesService = &purchasesService{}

func NewPurchasesService(_ context.Context, purchasesRepository repository.PurchasesRepository) service.PurchasesService {
	return &purchasesService{
		purchasesRepository: purchasesRepository,
	}
}

func (s *purchasesService) Save(ctx context.Context, input *dto.PurchasesCreateDto) (*dto.PurchasesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.purchasesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchasesService) UpdateByID(ctx context.Context, input *dto.PurchasesUpdateDto, pkRecordID int) (*dto.PurchasesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.purchasesRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchasesService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.purchasesRepository.DeleteByID(ctx, pkRecordID)
}

func (s *purchasesService) FindByID(ctx context.Context, pkRecordID int) (*dto.PurchasesDto, error) {
	entityById, err := s.purchasesRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchasesService) FindAll(ctx context.Context) ([]dto.PurchasesDto, error) {
	entities, err := s.purchasesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *purchasesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchasesDto, pageable.Page, error) {
	entities, page, err := s.purchasesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.PurchasesCreateDto) (*dbModel.Purchases, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchasesCreateDto->Purchases: input dto cannot be nil")
	}
	return &dbModel.Purchases{
		CustomerID:  input.CustomerID,
		Description: input.Description,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.PurchasesUpdateDto) (*dbModel.Purchases, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchasesUpdateDto->Purchases: input dto cannot be nil")
	}
	return &dbModel.Purchases{
		CustomerID:  input.CustomerID,
		Description: input.Description,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Purchases) ([]dto.PurchasesDto, error) {
	var outputDtos []dto.PurchasesDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Purchases) (dto.PurchasesDto, error) {
	if inputEntity == nil {
		return dto.PurchasesDto{}, fmt.Errorf("unexpected nil input for mapping between Purchases->PurchasesDto")
	}
	return dto.PurchasesDto{
		RecordID:    inputEntity.RecordID,
		CustomerID:  inputEntity.CustomerID,
		Description: inputEntity.Description,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		Guid:        inputEntity.Guid,
	}, nil
}
