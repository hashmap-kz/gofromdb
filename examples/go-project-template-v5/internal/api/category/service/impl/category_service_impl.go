package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/category/dto"
	dbModel "go-project-template-v5/internal/api/category/entity/postgres"

	"go-project-template-v5/internal/api/category/repository"
	"go-project-template-v5/internal/api/category/service"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

var _ service.CategoryService = &categoryService{}

func NewCategoryService(_ context.Context, categoryRepository repository.CategoryRepository) service.CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) Save(ctx context.Context, input *dto.CategoryCreateDto) (*dto.CategoryDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.categoryRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *categoryService) UpdateByID(ctx context.Context, entityId int, input *dto.CategoryUpdateDto) (*dto.CategoryDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.categoryRepository.UpdateByID(ctx, entityId, entityToUpdate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *categoryService) DeleteByID(ctx context.Context, id int) error {
	return s.categoryRepository.DeleteByID(ctx, id)
}

func (s *categoryService) FindByID(ctx context.Context, id int) (*dto.CategoryDto, error) {
	entityById, err := s.categoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *categoryService) FindAll(ctx context.Context) ([]dto.CategoryDto, error) {
	entities, err := s.categoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *categoryService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CategoryDto, pageable.Page, error) {
	entities, page, err := s.categoryRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.CategoryCreateDto) (*dbModel.Category, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CategoryCreateDto->Category: input dto cannot be nil")
	}
	return &dbModel.Category{
		Name:     input.Name,
		ParentID: input.ParentID,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.CategoryUpdateDto) (*dbModel.Category, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CategoryUpdateDto->Category: input dto cannot be nil")
	}
	return &dbModel.Category{
		Name:     input.Name,
		ParentID: input.ParentID,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Category) ([]dto.CategoryDto, error) {
	var outputDtos []dto.CategoryDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Category) (dto.CategoryDto, error) {
	if inputEntity == nil {
		return dto.CategoryDto{}, fmt.Errorf("unexpected nil input for mapping between Category->CategoryDto")
	}
	return dto.CategoryDto{
		RecordID:  inputEntity.RecordID,
		Name:      inputEntity.Name,
		ParentID:  inputEntity.ParentID,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}
