package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/composite_pk/dto"
	dbModel "go-project-template-v5/internal/api/composite_pk/entity/postgres"

	"go-project-template-v5/internal/api/composite_pk/repository"
	"go-project-template-v5/internal/api/composite_pk/service"
)

type compositePkService struct {
	compositePkRepository repository.CompositePkRepository
}

var _ service.CompositePkService = &compositePkService{}

func NewCompositePkService(_ context.Context, compositePkRepository repository.CompositePkRepository) service.CompositePkService {
	return &compositePkService{
		compositePkRepository: compositePkRepository,
	}
}

func (s *compositePkService) Save(ctx context.Context, input *dto.CompositePkCreateDto) (*dto.CompositePkDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.compositePkRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *compositePkService) UpdateByID(ctx context.Context, input *dto.CompositePkUpdateDto, pkTenantID int64, pkCode string) (*dto.CompositePkDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.compositePkRepository.UpdateByID(ctx, entityToUpdate, pkTenantID, pkCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *compositePkService) DeleteByID(ctx context.Context, pkTenantID int64, pkCode string) error {
	return s.compositePkRepository.DeleteByID(ctx, pkTenantID, pkCode)
}

func (s *compositePkService) FindByID(ctx context.Context, pkTenantID int64, pkCode string) (*dto.CompositePkDto, error) {
	entityByID, err := s.compositePkRepository.FindByID(ctx, pkTenantID, pkCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *compositePkService) FindAll(ctx context.Context) ([]dto.CompositePkDto, error) {
	entities, err := s.compositePkRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *compositePkService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CompositePkDto, pageable.Page, error) {
	entities, page, err := s.compositePkRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CompositePkCreateDto) (*dbModel.CompositePk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CompositePkCreateDto->CompositePk: input dto cannot be nil")
	}
	return &dbModel.CompositePk{
		TenantID: input.TenantID,
		Code:     input.Code,
		Name:     input.Name,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CompositePkUpdateDto) (*dbModel.CompositePk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CompositePkUpdateDto->CompositePk: input dto cannot be nil")
	}
	return &dbModel.CompositePk{
		Name: input.Name,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.CompositePk) ([]dto.CompositePkDto, error) {
	outputDtos := make([]dto.CompositePkDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.CompositePk) (dto.CompositePkDto, error) {
	if inputEntity == nil {
		return dto.CompositePkDto{}, fmt.Errorf("unexpected nil input for mapping between CompositePk->CompositePkDto")
	}
	return dto.CompositePkDto{
		TenantID: inputEntity.TenantID,
		Code:     inputEntity.Code,
		Name:     inputEntity.Name,
	}, nil
}
