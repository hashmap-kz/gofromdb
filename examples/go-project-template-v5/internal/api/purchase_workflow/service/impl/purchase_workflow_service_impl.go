package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_workflow/dto"
	dbModel "go-project-template-v5/internal/api/purchase_workflow/entity/postgres"

	"go-project-template-v5/internal/api/purchase_workflow/repository"
	"go-project-template-v5/internal/api/purchase_workflow/service"
)

type purchaseWorkflowService struct {
	purchaseWorkflowRepository repository.PurchaseWorkflowRepository
}

var _ service.PurchaseWorkflowService = &purchaseWorkflowService{}

func NewPurchaseWorkflowService(_ context.Context, purchaseWorkflowRepository repository.PurchaseWorkflowRepository) service.PurchaseWorkflowService {
	return &purchaseWorkflowService{
		purchaseWorkflowRepository: purchaseWorkflowRepository,
	}
}

func (s *purchaseWorkflowService) Save(ctx context.Context, input *dto.PurchaseWorkflowCreateDto) (*dto.PurchaseWorkflowDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.purchaseWorkflowRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseWorkflowService) UpdateByID(ctx context.Context, input *dto.PurchaseWorkflowUpdateDto, pkRecordID int) (*dto.PurchaseWorkflowDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.purchaseWorkflowRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseWorkflowService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.purchaseWorkflowRepository.DeleteByID(ctx, pkRecordID)
}

func (s *purchaseWorkflowService) FindByID(ctx context.Context, pkRecordID int) (*dto.PurchaseWorkflowDto, error) {
	entityById, err := s.purchaseWorkflowRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseWorkflowService) FindAll(ctx context.Context) ([]dto.PurchaseWorkflowDto, error) {
	entities, err := s.purchaseWorkflowRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *purchaseWorkflowService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchaseWorkflowDto, pageable.Page, error) {
	entities, page, err := s.purchaseWorkflowRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.PurchaseWorkflowCreateDto) (*dbModel.PurchaseWorkflow, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchaseWorkflowCreateDto->PurchaseWorkflow: input dto cannot be nil")
	}
	return &dbModel.PurchaseWorkflow{
		ValidPeriod:    input.ValidPeriod,
		BuyID:          input.BuyID,
		PurchaseStepID: input.PurchaseStepID,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.PurchaseWorkflowUpdateDto) (*dbModel.PurchaseWorkflow, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchaseWorkflowUpdateDto->PurchaseWorkflow: input dto cannot be nil")
	}
	return &dbModel.PurchaseWorkflow{
		ValidPeriod:    input.ValidPeriod,
		BuyID:          input.BuyID,
		PurchaseStepID: input.PurchaseStepID,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.PurchaseWorkflow) ([]dto.PurchaseWorkflowDto, error) {
	var outputDtos []dto.PurchaseWorkflowDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.PurchaseWorkflow) (dto.PurchaseWorkflowDto, error) {
	if inputEntity == nil {
		return dto.PurchaseWorkflowDto{}, fmt.Errorf("unexpected nil input for mapping between PurchaseWorkflow->PurchaseWorkflowDto")
	}
	return dto.PurchaseWorkflowDto{
		RecordID:       inputEntity.RecordID,
		ValidPeriod:    inputEntity.ValidPeriod,
		BuyID:          inputEntity.BuyID,
		PurchaseStepID: inputEntity.PurchaseStepID,
		CreatedAt:      inputEntity.CreatedAt,
		UpdatedAt:      inputEntity.UpdatedAt,
		Guid:           inputEntity.Guid,
	}, nil
}
