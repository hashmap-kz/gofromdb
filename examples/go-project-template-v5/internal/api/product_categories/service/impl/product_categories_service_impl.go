package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product_categories/dto"
	dbModel "go-project-template-v5/internal/api/product_categories/entity/postgres"

	"go-project-template-v5/internal/api/product_categories/repository"
	"go-project-template-v5/internal/api/product_categories/service"
)

type productCategoriesService struct {
	productCategoriesRepository repository.ProductCategoriesRepository
}

var _ service.ProductCategoriesService = &productCategoriesService{}

func NewProductCategoriesService(_ context.Context, productCategoriesRepository repository.ProductCategoriesRepository) service.ProductCategoriesService {
	return &productCategoriesService{
		productCategoriesRepository: productCategoriesRepository,
	}
}

func (s *productCategoriesService) Save(ctx context.Context, input *dto.ProductCategoriesCreateDto) (*dto.ProductCategoriesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.productCategoriesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productCategoriesService) UpdateByID(ctx context.Context, input *dto.ProductCategoriesUpdateDto, pkRecordID int) (*dto.ProductCategoriesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.productCategoriesRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productCategoriesService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.productCategoriesRepository.DeleteByID(ctx, pkRecordID)
}

func (s *productCategoriesService) FindByID(ctx context.Context, pkRecordID int) (*dto.ProductCategoriesDto, error) {
	entityById, err := s.productCategoriesRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productCategoriesService) FindAll(ctx context.Context) ([]dto.ProductCategoriesDto, error) {
	entities, err := s.productCategoriesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *productCategoriesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductCategoriesDto, pageable.Page, error) {
	entities, page, err := s.productCategoriesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.ProductCategoriesCreateDto) (*dbModel.ProductCategories, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductCategoriesCreateDto->ProductCategories: input dto cannot be nil")
	}
	return &dbModel.ProductCategories{
		Name:        input.Name,
		ParentID:    input.ParentID,
		ValidPeriod: input.ValidPeriod,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.ProductCategoriesUpdateDto) (*dbModel.ProductCategories, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductCategoriesUpdateDto->ProductCategories: input dto cannot be nil")
	}
	return &dbModel.ProductCategories{
		Name:        input.Name,
		ParentID:    input.ParentID,
		ValidPeriod: input.ValidPeriod,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.ProductCategories) ([]dto.ProductCategoriesDto, error) {
	var outputDtos []dto.ProductCategoriesDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.ProductCategories) (dto.ProductCategoriesDto, error) {
	if inputEntity == nil {
		return dto.ProductCategoriesDto{}, fmt.Errorf("unexpected nil input for mapping between ProductCategories->ProductCategoriesDto")
	}
	return dto.ProductCategoriesDto{
		RecordID:    inputEntity.RecordID,
		Name:        inputEntity.Name,
		ParentID:    inputEntity.ParentID,
		ValidPeriod: inputEntity.ValidPeriod,
		IsCurrent:   inputEntity.IsCurrent,
		CreatedAt:   inputEntity.CreatedAt,
		UpdatedAt:   inputEntity.UpdatedAt,
		Guid:        inputEntity.Guid,
	}, nil
}
