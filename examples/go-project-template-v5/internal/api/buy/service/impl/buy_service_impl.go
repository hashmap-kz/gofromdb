package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/buy/dto"
	dbModel "go-project-template-v5/internal/api/buy/entity/postgres"

	"go-project-template-v5/internal/api/buy/repository"
	"go-project-template-v5/internal/api/buy/service"
)

type buyService struct {
	repo repository.BuyRepository
}

var _ service.BuyService = &buyService{}

func NewBuyService(_ context.Context, repo repository.BuyRepository) service.BuyService {
	return &buyService{repo: repo}
}

func (s *buyService) Save(ctx context.Context, input *dto.BuyCreateDto) (*dto.BuyDto, error) {
	save, err := s.repo.Save(ctx, &dbModel.Buy{
		ClientID:    input.ClientID,
		Description: input.Description,
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

func (s *buyService) GetAll(ctx context.Context) ([]dto.BuyDto, error) {
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

func (s *buyService) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyDto, pageable.Page, error) {
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

func (s *buyService) Update(ctx context.Context, input *dto.BuyUpdateDto) (*dto.BuyDto, error) {
	// update dbModel
	updatedResult, err := s.repo.Update(ctx, &dbModel.Buy{
		RecordID:    input.RecordID,
		ClientID:    input.ClientID,
		Description: input.Description,
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

func (s *buyService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *buyService) GetByID(ctx context.Context, id int) (*dto.BuyDto, error) {
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

func mapEntitiesToDtos(inputEntities []dbModel.Buy) ([]dto.BuyDto, error) {
	var outputDtos []dto.BuyDto
	for _, inputEntity := range inputEntities {
		toDto, err := mapEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func mapEntityToDto(inputEntity *dbModel.Buy) (dto.BuyDto, error) {
	if inputEntity == nil {
		return dto.BuyDto{}, fmt.Errorf("unexpected nil input for mapping between Buy->BuyDto")
	}
	return dto.BuyDto{
		RecordID:    inputEntity.RecordID,
		ClientID:    inputEntity.ClientID,
		Description: inputEntity.Description,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		Guid:        inputEntity.Guid,
	}, nil
}
