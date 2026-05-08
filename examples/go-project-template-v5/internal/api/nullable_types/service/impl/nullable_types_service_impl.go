package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/nullable_types/dto"
	dbModel "go-project-template-v5/internal/api/nullable_types/entity/postgres"

	"go-project-template-v5/internal/api/nullable_types/repository"
	"go-project-template-v5/internal/api/nullable_types/service"
)

type nullableTypesService struct {
	nullableTypesRepository repository.NullableTypesRepository
}

var _ service.NullableTypesService = &nullableTypesService{}

func NewNullableTypesService(_ context.Context, nullableTypesRepository repository.NullableTypesRepository) service.NullableTypesService {
	return &nullableTypesService{
		nullableTypesRepository: nullableTypesRepository,
	}
}

func (s *nullableTypesService) Save(ctx context.Context, input *dto.NullableTypesCreateDto) (*dto.NullableTypesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.nullableTypesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *nullableTypesService) UpdateByID(ctx context.Context, input *dto.NullableTypesUpdateDto, pkID int64) (*dto.NullableTypesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.nullableTypesRepository.UpdateByID(ctx, entityToUpdate, pkID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *nullableTypesService) DeleteByID(ctx context.Context, pkID int64) error {
	return s.nullableTypesRepository.DeleteByID(ctx, pkID)
}

func (s *nullableTypesService) FindByID(ctx context.Context, pkID int64) (*dto.NullableTypesDto, error) {
	entityByID, err := s.nullableTypesRepository.FindByID(ctx, pkID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *nullableTypesService) FindAll(ctx context.Context) ([]dto.NullableTypesDto, error) {
	entities, err := s.nullableTypesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *nullableTypesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.NullableTypesDto, pageable.Page, error) {
	entities, page, err := s.nullableTypesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.NullableTypesCreateDto) (*dbModel.NullableTypes, error) {
	if input == nil {
		return nil, fmt.Errorf("convert NullableTypesCreateDto->NullableTypes: input dto cannot be nil")
	}
	return &dbModel.NullableTypes{
		Name:    input.Name,
		Amount:  input.Amount,
		Payload: input.Payload,
		Tags:    input.Tags,
		Active:  input.Active,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.NullableTypesUpdateDto) (*dbModel.NullableTypes, error) {
	if input == nil {
		return nil, fmt.Errorf("convert NullableTypesUpdateDto->NullableTypes: input dto cannot be nil")
	}
	return &dbModel.NullableTypes{
		Name:    input.Name,
		Amount:  input.Amount,
		Payload: input.Payload,
		Tags:    input.Tags,
		Active:  input.Active,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.NullableTypes) ([]dto.NullableTypesDto, error) {
	outputDtos := make([]dto.NullableTypesDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.NullableTypes) (dto.NullableTypesDto, error) {
	if inputEntity == nil {
		return dto.NullableTypesDto{}, fmt.Errorf("unexpected nil input for mapping between NullableTypes->NullableTypesDto")
	}
	return dto.NullableTypesDto{
		ID:        inputEntity.ID,
		Name:      inputEntity.Name,
		Amount:    inputEntity.Amount,
		Payload:   inputEntity.Payload,
		Tags:      inputEntity.Tags,
		Active:    inputEntity.Active,
		CreatedAt: inputEntity.CreatedAt,
	}, nil
}
