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
	buyRepository repository.BuyRepository
}

var _ service.BuyService = &buyService{}

func NewBuyService(_ context.Context, buyRepository repository.BuyRepository) service.BuyService {
	return &buyService{
		buyRepository: buyRepository,
	}
}

func (s *buyService) Save(ctx context.Context, input *dto.BuyCreateDto) (*dto.BuyDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.buyRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyService) UpdateByID(ctx context.Context, pkRecordID int, input *dto.BuyUpdateDto) (*dto.BuyDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.buyRepository.UpdateByID(ctx, pkRecordID, entityToUpdate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.buyRepository.DeleteByID(ctx, pkRecordID)
}

func (s *buyService) FindByID(ctx context.Context, pkRecordID int) (*dto.BuyDto, error) {
	entityById, err := s.buyRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *buyService) FindAll(ctx context.Context) ([]dto.BuyDto, error) {
	entities, err := s.buyRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *buyService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.BuyDto, pageable.Page, error) {
	entities, page, err := s.buyRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.BuyCreateDto) (*dbModel.Buy, error) {
	if input == nil {
		return nil, fmt.Errorf("convert BuyCreateDto->Buy: input dto cannot be nil")
	}
	return &dbModel.Buy{
		ClientID:    input.ClientID,
		Description: input.Description,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.BuyUpdateDto) (*dbModel.Buy, error) {
	if input == nil {
		return nil, fmt.Errorf("convert BuyUpdateDto->Buy: input dto cannot be nil")
	}
	return &dbModel.Buy{
		ClientID:    input.ClientID,
		Description: input.Description,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Buy) ([]dto.BuyDto, error) {
	var outputDtos []dto.BuyDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Buy) (dto.BuyDto, error) {
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
