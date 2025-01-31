package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy_item/dto"
	dbModel "go-project-template-v5/internal/api/buy_item/entity/postgres"

	"go-project-template-v5/internal/api/buy_item/repository"
	"go-project-template-v5/internal/api/buy_item/service"
)

type buyItemService struct {
	buyItemRepository repository.BuyItemRepository
}

var _ service.BuyItemService = &buyItemService{}

func NewBuyItemService(_ context.Context, buyItemRepository repository.BuyItemRepository) service.BuyItemService {
	return &buyItemService{
		buyItemRepository: buyItemRepository,
	}
}

func (s *buyItemService) Save(ctx context.Context, input *dto.BuyItemCreateDto) (*dto.BuyItemDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.buyItemRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyItemService) UpdateByID(ctx context.Context, entityId int, input *dto.BuyItemUpdateDto) (*dto.BuyItemDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.buyItemRepository.UpdateByID(ctx, entityId, entityToUpdate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyItemService) DeleteByID(ctx context.Context, id int) error {
	return s.buyItemRepository.DeleteByID(ctx, id)
}

func (s *buyItemService) FindByID(ctx context.Context, id int) (*dto.BuyItemDto, error) {
	entityById, err := s.buyItemRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyItemService) FindAll(ctx context.Context) ([]dto.BuyItemDto, error) {
	entities, err := s.buyItemRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *buyItemService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyItemDto, pageable.Page, error) {
	entities, page, err := s.buyItemRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.BuyItemCreateDto) (*dbModel.BuyItem, error) {
	if input == nil {
		return nil, fmt.Errorf("convert BuyItemCreateDto->BuyItem: input dto cannot be nil")
	}
	return &dbModel.BuyItem{
		BuyID:     input.BuyID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.BuyItemUpdateDto) (*dbModel.BuyItem, error) {
	if input == nil {
		return nil, fmt.Errorf("convert BuyItemUpdateDto->BuyItem: input dto cannot be nil")
	}
	return &dbModel.BuyItem{
		BuyID:     input.BuyID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.BuyItem) ([]dto.BuyItemDto, error) {
	var outputDtos []dto.BuyItemDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.BuyItem) (dto.BuyItemDto, error) {
	if inputEntity == nil {
		return dto.BuyItemDto{}, fmt.Errorf("unexpected nil input for mapping between BuyItem->BuyItemDto")
	}
	return dto.BuyItemDto{
		RecordID:  inputEntity.RecordID,
		BuyID:     inputEntity.BuyID,
		ProductID: inputEntity.ProductID,
		Quantity:  inputEntity.Quantity,
		Price:     inputEntity.Price,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}
