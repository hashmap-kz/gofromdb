package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product/dto"
	dbModel "go-project-template-v5/internal/api/product/entity/postgres"

	"go-project-template-v5/internal/api/product/repository"
	"go-project-template-v5/internal/api/product/service"
)

type productService struct {
	repo repository.ProductRepository
}

var _ service.ProductService = &productService{}

func NewProductService(_ context.Context, repo repository.ProductRepository) service.ProductService {
	return &productService{repo: repo}
}

func (s *productService) Save(ctx context.Context, input *dto.ProductCreateDto) (*dto.ProductDto, error) {
	save, err := s.repo.Save(ctx, &dbModel.Product{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
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

func (s *productService) GetAll(ctx context.Context) ([]dto.ProductDto, error) {
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

func (s *productService) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductDto, pageable.Page, error) {
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

func (s *productService) Update(ctx context.Context, input *dto.ProductUpdateDto) (*dto.ProductDto, error) {
	// update dbModel
	updatedResult, err := s.repo.Update(ctx, &dbModel.Product{
		RecordID:    input.RecordID,
		CategoryID:  input.CategoryID,
		Name:        input.Name,
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

func (s *productService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *productService) GetByID(ctx context.Context, id int) (*dto.ProductDto, error) {
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

func mapEntitiesToDtos(inputEntities []dbModel.Product) ([]dto.ProductDto, error) {
	var outputDtos []dto.ProductDto
	for _, inputEntity := range inputEntities {
		toDto, err := mapEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func mapEntityToDto(inputEntity *dbModel.Product) (dto.ProductDto, error) {
	if inputEntity == nil {
		return dto.ProductDto{}, fmt.Errorf("unexpected nil input for mapping between Product->ProductDto")
	}
	return dto.ProductDto{
		RecordID:    inputEntity.RecordID,
		CategoryID:  inputEntity.CategoryID,
		Name:        inputEntity.Name,
		Description: inputEntity.Description,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		Guid:        inputEntity.Guid,
	}, nil
}
