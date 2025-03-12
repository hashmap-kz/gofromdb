package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/categories/dto"
	dbModel "go-project-template-v5/internal/api/categories/entity/postgres"

	"go-project-template-v5/internal/api/categories/repository"
	"go-project-template-v5/internal/api/categories/service"
)

type categoriesService struct {
	categoriesRepository repository.CategoriesRepository
}

var _ service.CategoriesService = &categoriesService{}

func NewCategoriesService(_ context.Context, categoriesRepository repository.CategoriesRepository) service.CategoriesService {
	return &categoriesService{
		categoriesRepository: categoriesRepository,
	}
}

func (s *categoriesService) Save(ctx context.Context, input *dto.CategoriesCreateDto) (*dto.CategoriesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.categoriesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *categoriesService) UpdateByID(ctx context.Context, input *dto.CategoriesUpdateDto, pkRecordID int) (*dto.CategoriesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.categoriesRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *categoriesService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.categoriesRepository.DeleteByID(ctx, pkRecordID)
}

func (s *categoriesService) FindByID(ctx context.Context, pkRecordID int) (*dto.CategoriesDto, error) {
	entityByID, err := s.categoriesRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *categoriesService) FindAll(ctx context.Context) ([]dto.CategoriesDto, error) {
	entities, err := s.categoriesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *categoriesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CategoriesDto, pageable.Page, error) {
	entities, page, err := s.categoriesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CategoriesCreateDto) (*dbModel.Categories, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CategoriesCreateDto->Categories: input dto cannot be nil")
	}
	return &dbModel.Categories{
		Name:        input.Name,
		ParentID:    input.ParentID,
		ValidPeriod: input.ValidPeriod,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CategoriesUpdateDto) (*dbModel.Categories, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CategoriesUpdateDto->Categories: input dto cannot be nil")
	}
	return &dbModel.Categories{
		Name:        input.Name,
		ParentID:    input.ParentID,
		ValidPeriod: input.ValidPeriod,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Categories) ([]dto.CategoriesDto, error) {
	outputDtos := make([]CategoriesDto, 0, len(inputEntities))
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Categories) (dto.CategoriesDto, error) {
	if inputEntity == nil {
		return dto.CategoriesDto{}, fmt.Errorf("unexpected nil input for mapping between Categories->CategoriesDto")
	}
	return dto.CategoriesDto{
		RecordID:    inputEntity.RecordID,
		Name:        inputEntity.Name,
		ParentID:    inputEntity.ParentID,
		ValidPeriod: inputEntity.ValidPeriod,
		IsCurrent:   inputEntity.IsCurrent,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		GUID:        inputEntity.GUID,
	}, nil
}
