package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/currencies/dto"
	dbModel "go-project-template-v5/internal/api/currencies/entity/postgres"

	"go-project-template-v5/internal/api/currencies/repository"
	"go-project-template-v5/internal/api/currencies/service"
)

type currenciesService struct {
	currenciesRepository repository.CurrenciesRepository
}

var _ service.CurrenciesService = &currenciesService{}

func NewCurrenciesService(_ context.Context, currenciesRepository repository.CurrenciesRepository) service.CurrenciesService {
	return &currenciesService{
		currenciesRepository: currenciesRepository,
	}
}

func (s *currenciesService) Save(ctx context.Context, input *dto.CurrenciesCreateDto) (*dto.CurrenciesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.currenciesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *currenciesService) UpdateByID(ctx context.Context, input *dto.CurrenciesUpdateDto, pkRecordID int) (*dto.CurrenciesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.currenciesRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *currenciesService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.currenciesRepository.DeleteByID(ctx, pkRecordID)
}

func (s *currenciesService) FindByID(ctx context.Context, pkRecordID int) (*dto.CurrenciesDto, error) {
	entityById, err := s.currenciesRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *currenciesService) FindAll(ctx context.Context) ([]dto.CurrenciesDto, error) {
	entities, err := s.currenciesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *currenciesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CurrenciesDto, pageable.Page, error) {
	entities, page, err := s.currenciesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CurrenciesCreateDto) (*dbModel.Currencies, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CurrenciesCreateDto->Currencies: input dto cannot be nil")
	}
	return &dbModel.Currencies{
		CurrencyCode:  input.CurrencyCode,
		CurrencyValue: input.CurrencyValue,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CurrenciesUpdateDto) (*dbModel.Currencies, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CurrenciesUpdateDto->Currencies: input dto cannot be nil")
	}
	return &dbModel.Currencies{
		CurrencyCode:  input.CurrencyCode,
		CurrencyValue: input.CurrencyValue,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Currencies) ([]dto.CurrenciesDto, error) {
	var outputDtos []dto.CurrenciesDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Currencies) (dto.CurrenciesDto, error) {
	if inputEntity == nil {
		return dto.CurrenciesDto{}, fmt.Errorf("unexpected nil input for mapping between Currencies->CurrenciesDto")
	}
	return dto.CurrenciesDto{
		RecordID:      inputEntity.RecordID,
		CurrencyCode:  inputEntity.CurrencyCode,
		CurrencyValue: inputEntity.CurrencyValue,
		CreatedAt:     inputEntity.CreatedAt,
		UpdatedAt:     inputEntity.UpdatedAt,
		Guid:          inputEntity.Guid,
	}, nil
}
