package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/purchase_items/dto"
	dbModel "go-project-template-v5/internal/api/purchase_items/entity/postgres"

	"go-project-template-v5/internal/api/purchase_items/repository"
	"go-project-template-v5/internal/api/purchase_items/service"
)

type purchaseItemsService struct {
	purchaseItemsRepository repository.PurchaseItemsRepository
}

var _ service.PurchaseItemsService = &purchaseItemsService{}

func NewPurchaseItemsService(_ context.Context, purchaseItemsRepository repository.PurchaseItemsRepository) service.PurchaseItemsService {
	return &purchaseItemsService{
		purchaseItemsRepository: purchaseItemsRepository,
	}
}

func (s *purchaseItemsService) Save(ctx context.Context, input *dto.PurchaseItemsCreateDto) (*dto.PurchaseItemsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.purchaseItemsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseItemsService) UpdateByID(ctx context.Context, input *dto.PurchaseItemsUpdateDto, pkRecordID int) (*dto.PurchaseItemsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.purchaseItemsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseItemsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.purchaseItemsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *purchaseItemsService) FindByID(ctx context.Context, pkRecordID int) (*dto.PurchaseItemsDto, error) {
	entityById, err := s.purchaseItemsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *purchaseItemsService) FindAll(ctx context.Context) ([]dto.PurchaseItemsDto, error) {
	entities, err := s.purchaseItemsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *purchaseItemsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.PurchaseItemsDto, pageable.Page, error) {
	entities, page, err := s.purchaseItemsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.PurchaseItemsCreateDto) (*dbModel.PurchaseItems, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchaseItemsCreateDto->PurchaseItems: input dto cannot be nil")
	}
	return &dbModel.PurchaseItems{
		PurchaseID: input.PurchaseID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		Price:      input.Price,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.PurchaseItemsUpdateDto) (*dbModel.PurchaseItems, error) {
	if input == nil {
		return nil, fmt.Errorf("convert PurchaseItemsUpdateDto->PurchaseItems: input dto cannot be nil")
	}
	return &dbModel.PurchaseItems{
		PurchaseID: input.PurchaseID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		Price:      input.Price,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.PurchaseItems) ([]dto.PurchaseItemsDto, error) {
	var outputDtos []dto.PurchaseItemsDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.PurchaseItems) (dto.PurchaseItemsDto, error) {
	if inputEntity == nil {
		return dto.PurchaseItemsDto{}, fmt.Errorf("unexpected nil input for mapping between PurchaseItems->PurchaseItemsDto")
	}
	return dto.PurchaseItemsDto{
		RecordID:   inputEntity.RecordID,
		PurchaseID: inputEntity.PurchaseID,
		ProductID:  inputEntity.ProductID,
		Quantity:   inputEntity.Quantity,
		Price:      inputEntity.Price,
		CreatedAt:  inputEntity.CreatedAt,
		UpdatedAt:  inputEntity.UpdatedAt,
		Guid:       inputEntity.Guid,
	}, nil
}
