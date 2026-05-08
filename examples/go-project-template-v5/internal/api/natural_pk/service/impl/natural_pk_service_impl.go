package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/natural_pk/dto"
	dbModel "go-project-template-v5/internal/api/natural_pk/entity/postgres"

	"go-project-template-v5/internal/api/natural_pk/repository"
	"go-project-template-v5/internal/api/natural_pk/service"
)

type naturalPkService struct {
	naturalPkRepository repository.NaturalPkRepository
}

var _ service.NaturalPkService = &naturalPkService{}

func NewNaturalPkService(_ context.Context, naturalPkRepository repository.NaturalPkRepository) service.NaturalPkService {
	return &naturalPkService{
		naturalPkRepository: naturalPkRepository,
	}
}

func (s *naturalPkService) Save(ctx context.Context, input *dto.NaturalPkCreateDto) (*dto.NaturalPkDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.naturalPkRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *naturalPkService) UpdateByID(ctx context.Context, input *dto.NaturalPkUpdateDto, pkCode string) (*dto.NaturalPkDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.naturalPkRepository.UpdateByID(ctx, entityToUpdate, pkCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *naturalPkService) DeleteByID(ctx context.Context, pkCode string) error {
	return s.naturalPkRepository.DeleteByID(ctx, pkCode)
}

func (s *naturalPkService) FindByID(ctx context.Context, pkCode string) (*dto.NaturalPkDto, error) {
	entityByID, err := s.naturalPkRepository.FindByID(ctx, pkCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *naturalPkService) FindAll(ctx context.Context) ([]dto.NaturalPkDto, error) {
	entities, err := s.naturalPkRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *naturalPkService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.NaturalPkDto, pageable.Page, error) {
	entities, page, err := s.naturalPkRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.NaturalPkCreateDto) (*dbModel.NaturalPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert NaturalPkCreateDto->NaturalPk: input dto cannot be nil")
	}
	return &dbModel.NaturalPk{
		Code: input.Code,
		Name: input.Name,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.NaturalPkUpdateDto) (*dbModel.NaturalPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert NaturalPkUpdateDto->NaturalPk: input dto cannot be nil")
	}
	return &dbModel.NaturalPk{
		Name: input.Name,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.NaturalPk) ([]dto.NaturalPkDto, error) {
	outputDtos := make([]dto.NaturalPkDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.NaturalPk) (dto.NaturalPkDto, error) {
	if inputEntity == nil {
		return dto.NaturalPkDto{}, fmt.Errorf("unexpected nil input for mapping between NaturalPk->NaturalPkDto")
	}
	return dto.NaturalPkDto{
		Code: inputEntity.Code,
		Name: inputEntity.Name,
	}, nil
}
