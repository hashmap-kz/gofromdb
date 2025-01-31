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
	productRepository repository.ProductRepository
}

var _ service.ProductService = &productService{}

func NewProductService(_ context.Context, productRepository repository.ProductRepository) service.ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) Save(ctx context.Context, input *dto.ProductCreateDto) (*dto.ProductDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.productRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productService) UpdateByID(ctx context.Context, entityId int, input *dto.ProductUpdateDto) (*dto.ProductDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.productRepository.UpdateByID(ctx, entityId, entityToUpdate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productService) DeleteByID(ctx context.Context, id int) error {
	return s.productRepository.DeleteByID(ctx, id)
}

func (s *productService) FindByID(ctx context.Context, id int) (*dto.ProductDto, error) {
	entityById, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productService) FindAll(ctx context.Context) ([]dto.ProductDto, error) {
	entities, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *productService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductDto, pageable.Page, error) {
	entities, page, err := s.productRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.ProductCreateDto) (*dbModel.Product, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductCreateDto->Product: input dto cannot be nil")
	}
	return &dbModel.Product{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Description: input.Description,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.ProductUpdateDto) (*dbModel.Product, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductUpdateDto->Product: input dto cannot be nil")
	}
	return &dbModel.Product{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Description: input.Description,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Product) ([]dto.ProductDto, error) {
	var outputDtos []dto.ProductDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Product) (dto.ProductDto, error) {
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
