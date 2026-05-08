package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/uuid_pk/dto"
	dbModel "go-project-template-v5/internal/api/uuid_pk/entity/postgres"

	"go-project-template-v5/internal/api/uuid_pk/repository"
	"go-project-template-v5/internal/api/uuid_pk/service"
)

type uUIDPkService struct {
	uUIDPkRepository repository.UUIDPkRepository
}

var _ service.UUIDPkService = &uUIDPkService{}

func NewUUIDPkService(_ context.Context, uUIDPkRepository repository.UUIDPkRepository) service.UUIDPkService {
	return &uUIDPkService{
		uUIDPkRepository: uUIDPkRepository,
	}
}

func (s *uUIDPkService) Save(ctx context.Context, input *dto.UUIDPkCreateDto) (*dto.UUIDPkDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.uUIDPkRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *uUIDPkService) UpdateByID(ctx context.Context, input *dto.UUIDPkUpdateDto, pkID string) (*dto.UUIDPkDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.uUIDPkRepository.UpdateByID(ctx, entityToUpdate, pkID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *uUIDPkService) DeleteByID(ctx context.Context, pkID string) error {
	return s.uUIDPkRepository.DeleteByID(ctx, pkID)
}

func (s *uUIDPkService) FindByID(ctx context.Context, pkID string) (*dto.UUIDPkDto, error) {
	entityByID, err := s.uUIDPkRepository.FindByID(ctx, pkID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *uUIDPkService) FindAll(ctx context.Context) ([]dto.UUIDPkDto, error) {
	entities, err := s.uUIDPkRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *uUIDPkService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.UUIDPkDto, pageable.Page, error) {
	entities, page, err := s.uUIDPkRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.UUIDPkCreateDto) (*dbModel.UUIDPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UUIDPkCreateDto->UUIDPk: input dto cannot be nil")
	}
	return &dbModel.UUIDPk{
		ID:   input.ID,
		Name: input.Name,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.UUIDPkUpdateDto) (*dbModel.UUIDPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UUIDPkUpdateDto->UUIDPk: input dto cannot be nil")
	}
	return &dbModel.UUIDPk{
		Name: input.Name,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.UUIDPk) ([]dto.UUIDPkDto, error) {
	outputDtos := make([]dto.UUIDPkDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.UUIDPk) (dto.UUIDPkDto, error) {
	if inputEntity == nil {
		return dto.UUIDPkDto{}, fmt.Errorf("unexpected nil input for mapping between UUIDPk->UUIDPkDto")
	}
	return dto.UUIDPkDto{
		ID:   inputEntity.ID,
		Name: inputEntity.Name,
	}, nil
}
