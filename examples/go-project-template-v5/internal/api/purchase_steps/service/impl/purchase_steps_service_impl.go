package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_steps/dto"
	dbModel "go-project-template-v5/internal/api/purchase_steps/entity/postgres"

	"go-project-template-v5/internal/api/purchase_steps/repository"
	"go-project-template-v5/internal/api/purchase_steps/service"
)

type purchaseStepsService struct {
	purchaseStepsRepository repository.PurchaseStepsRepository
}

var _ service.PurchaseStepsService = &purchaseStepsService{}

func NewPurchaseStepsService(_ context.Context, purchaseStepsRepository repository.PurchaseStepsRepository) service.PurchaseStepsService {
	return &purchaseStepsService{
		purchaseStepsRepository: purchaseStepsRepository,
	}
}

func (s *purchaseStepsService) Save(ctx context.Context, input *dto.PurchaseStepsCreateDto) (*dto.PurchaseStepsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.purchaseStepsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseStepsService) UpdateByID(ctx context.Context, input *dto.PurchaseStepsUpdateDto, pkRecordID int) (*dto.PurchaseStepsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.purchaseStepsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseStepsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.purchaseStepsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *purchaseStepsService) FindByID(ctx context.Context, pkRecordID int) (*dto.PurchaseStepsDto, error) {
	entityById, err := s.purchaseStepsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseStepsService) FindAll(ctx context.Context) ([]dto.PurchaseStepsDto, error) {
	entities, err := s.purchaseStepsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *purchaseStepsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchaseStepsDto, pageable.Page, error) {
	entities, page, err := s.purchaseStepsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.PurchaseStepsCreateDto) (*dbModel.PurchaseSteps, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchaseStepsCreateDto->PurchaseSteps: input dto cannot be nil")
	}
	return &dbModel.PurchaseSteps{
		ValidPeriod: input.ValidPeriod,
		BuyID:       input.BuyID,
		StepID:      input.StepID,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.PurchaseStepsUpdateDto) (*dbModel.PurchaseSteps, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchaseStepsUpdateDto->PurchaseSteps: input dto cannot be nil")
	}
	return &dbModel.PurchaseSteps{
		ValidPeriod: input.ValidPeriod,
		BuyID:       input.BuyID,
		StepID:      input.StepID,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.PurchaseSteps) ([]dto.PurchaseStepsDto, error) {
	var outputDtos []dto.PurchaseStepsDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.PurchaseSteps) (dto.PurchaseStepsDto, error) {
	if inputEntity == nil {
		return dto.PurchaseStepsDto{}, fmt.Errorf("unexpected nil input for mapping between PurchaseSteps->PurchaseStepsDto")
	}
	return dto.PurchaseStepsDto{
		RecordID:    inputEntity.RecordID,
		ValidPeriod: inputEntity.ValidPeriod,
		BuyID:       inputEntity.BuyID,
		StepID:      inputEntity.StepID,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		Guid:        inputEntity.Guid,
	}, nil
}
