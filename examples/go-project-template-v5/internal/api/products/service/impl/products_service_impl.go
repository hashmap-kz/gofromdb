package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/products/dto"
	dbModel "go-project-template-v5/internal/api/products/entity/postgres"

	"go-project-template-v5/internal/api/products/repository"
	"go-project-template-v5/internal/api/products/service"
)

type productsService struct {
	productsRepository repository.ProductsRepository
}

var _ service.ProductsService = &productsService{}

func NewProductsService(_ context.Context, productsRepository repository.ProductsRepository) service.ProductsService {
	return &productsService{
		productsRepository: productsRepository,
	}
}

func (s *productsService) Save(ctx context.Context, input *dto.ProductsCreateDto) (*dto.ProductsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.productsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productsService) UpdateByID(ctx context.Context, input *dto.ProductsUpdateDto, pkRecordID int) (*dto.ProductsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.productsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.productsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *productsService) FindByID(ctx context.Context, pkRecordID int) (*dto.ProductsDto, error) {
	entityByID, err := s.productsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productsService) FindAll(ctx context.Context) ([]dto.ProductsDto, error) {
	entities, err := s.productsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *productsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductsDto, pageable.Page, error) {
	entities, page, err := s.productsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.ProductsCreateDto) (*dbModel.Products, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductsCreateDto->Products: input dto cannot be nil")
	}
	return &dbModel.Products{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Description: input.Description,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.ProductsUpdateDto) (*dbModel.Products, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductsUpdateDto->Products: input dto cannot be nil")
	}
	return &dbModel.Products{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Description: input.Description,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Products) ([]dto.ProductsDto, error) {
	outputDtos := make([]ProductsDto, 0, len(inputEntities))
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Products) (dto.ProductsDto, error) {
	if inputEntity == nil {
		return dto.ProductsDto{}, fmt.Errorf("unexpected nil input for mapping between Products->ProductsDto")
	}
	return dto.ProductsDto{
		RecordID:    inputEntity.RecordID,
		CategoryID:  inputEntity.CategoryID,
		Name:        inputEntity.Name,
		Description: inputEntity.Description,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		GUID:        inputEntity.GUID,
	}, nil
}
