package impl

import (
	"context"
	"go-project-template-v5/internal/api/buy_item/dto"
	"go-project-template-v5/internal/api/buy_item/repository"
	"go-project-template-v5/internal/api/buy_item/service"
	"go-project-template-v5/pkg/pageable"
)

type buyItemService struct {
	repo repository.BuyItemRepository
}

var _ service.BuyItemService = &buyItemService{}

func NewBuyItemService(_ context.Context, repo repository.BuyItemRepository) service.BuyItemService {
	return &buyItemService{repo: repo}
}

func (s *buyItemService) Save(ctx context.Context, input *dto.BuyItemCreateDto) (*dto.BuyItemDto, error) {
	entity, err := dto.FromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.repo.Save(ctx, entity)
	if err != nil {
		return nil, err
	}
	toDto, err := dto.FromEntityToDto(save)
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
	toDtos, err := dto.FromEntitiesToDtos(entities)
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
	toDtos, err := dto.FromEntitiesToDtos(entities)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return toDtos, page, nil
}

func (s *buyItemService) Update(ctx context.Context, entityId int, input *dto.BuyItemUpdateDto) (*dto.BuyItemDto, error) {
	entity, err := dto.FromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.repo.Update(ctx, entityId, entity)
	if err != nil {
		return nil, err
	}
	toDto, err := dto.FromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyItemService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *buyItemService) GetByID(ctx context.Context, id int) (*dto.BuyItemDto, error) {
	entityById, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	toDto, err := dto.FromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}
