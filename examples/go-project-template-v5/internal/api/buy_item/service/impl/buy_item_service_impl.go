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
	repo repository.BuyItemRepository
}

var _ service.BuyItemService = &buyItemService{}

func NewBuyItemService(_ context.Context, repo repository.BuyItemRepository) service.BuyItemService {
	return &buyItemService{repo: repo}
}

func (s *buyItemService) Save(ctx context.Context, input *dto.BuyItemCreateDto) (*dto.BuyItemDto, error) {
	save, err := s.repo.Save(ctx, &dbModel.BuyItem{
		BuyID:     input.BuyID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	})
	if err != nil {
		return nil, err
	}
	toDto, err := mapEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyItemService) GetAll(ctx context.Context) ([]dto.BuyItemDto, error) {
	entities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := mapEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *buyItemService) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyItemDto, pageable.Page, error) {
	entities, page, err := s.repo.GetAllPaginated(ctx, pq)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	toDtos, err := mapEntitiesToDtos(entities)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return toDtos, page, nil
}

func (s *buyItemService) Update(ctx context.Context, input *dto.BuyItemUpdateDto) (*dto.BuyItemDto, error) {
	// update dbModel
	updatedResult, err := s.repo.Update(ctx, &dbModel.BuyItem{
		RecordID:  input.RecordID,
		BuyID:     input.BuyID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Price:     input.Price,
	})
	if err != nil {
		return nil, err
	}

	// convert dbModel to internal dto
	toDto, err := mapEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}

	return &toDto, err
}

func (s *buyItemService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *buyItemService) GetByID(ctx context.Context, id int) (*dto.BuyItemDto, error) {
	// retrieve dbModel by id
	entityById, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// convert dbModel to internal dto
	toDto, err := mapEntityToDto(entityById)
	if err != nil {
		return nil, err
	}

	return &toDto, err
}

// mappers

func mapEntitiesToDtos(inputEntities []dbModel.BuyItem) ([]dto.BuyItemDto, error) {
	var outputDtos []dto.BuyItemDto
	for _, inputEntity := range inputEntities {
		toDto, err := mapEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func mapEntityToDto(inputEntity *dbModel.BuyItem) (dto.BuyItemDto, error) {
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
