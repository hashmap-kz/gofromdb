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
	repo repository.CategoryRepository
}

var _ service.CategoryService = &categoryService{}

func NewCategoryService(_ context.Context, repo repository.CategoryRepository) service.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Save(ctx context.Context, input *dto.CategoryCreateDto) (*dto.CategoryDto, error) {
	save, err := s.repo.Save(ctx, &dbModel.Category{
		Name:     input.Name,
		ParentID: input.ParentID,
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

func (s *categoryService) GetAll(ctx context.Context) ([]dto.CategoryDto, error) {
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

func (s *categoryService) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.CategoryDto, pageable.Page, error) {
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

func (s *categoryService) Update(ctx context.Context, entityId int, input *dto.CategoryUpdateDto) (*dto.CategoryDto, error) {
	// update dbModel
	updatedResult, err := s.repo.Update(ctx, entityId, &dbModel.Category{
		Name:     input.Name,
		ParentID: input.ParentID,
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

func (s *categoryService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *categoryService) GetByID(ctx context.Context, id int) (*dto.CategoryDto, error) {
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

func mapEntitiesToDtos(inputEntities []dbModel.Category) ([]dto.CategoryDto, error) {
	var outputDtos []dto.CategoryDto
	for _, inputEntity := range inputEntities {
		toDto, err := mapEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func mapEntityToDto(inputEntity *dbModel.Category) (dto.CategoryDto, error) {
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
