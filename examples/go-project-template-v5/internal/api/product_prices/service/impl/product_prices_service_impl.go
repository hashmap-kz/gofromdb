package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/product_prices/dto"
	dbModel "go-project-template-v5/internal/api/product_prices/entity/postgres"

	"go-project-template-v5/internal/api/product_prices/repository"
	"go-project-template-v5/internal/api/product_prices/service"
)

type productPricesService struct {
	productPricesRepository repository.ProductPricesRepository
}

var _ service.ProductPricesService = &productPricesService{}

func NewProductPricesService(_ context.Context, productPricesRepository repository.ProductPricesRepository) service.ProductPricesService {
	return &productPricesService{
		productPricesRepository: productPricesRepository,
	}
}

func (s *productPricesService) Save(ctx context.Context, input *dto.ProductPricesCreateDto) (*dto.ProductPricesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.productPricesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productPricesService) UpdateByID(ctx context.Context, input *dto.ProductPricesUpdateDto, pkRecordID int) (*dto.ProductPricesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.productPricesRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productPricesService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.productPricesRepository.DeleteByID(ctx, pkRecordID)
}

func (s *productPricesService) FindByID(ctx context.Context, pkRecordID int) (*dto.ProductPricesDto, error) {
	entityById, err := s.productPricesRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *productPricesService) FindAll(ctx context.Context) ([]dto.ProductPricesDto, error) {
	entities, err := s.productPricesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *productPricesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ProductPricesDto, pageable.Page, error) {
	entities, page, err := s.productPricesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.ProductPricesCreateDto) (*dbModel.ProductPrices, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductPricesCreateDto->ProductPrices: input dto cannot be nil")
	}
	return &dbModel.ProductPrices{
		ProductPricePeriod: input.ProductPricePeriod,
		ProductID:          input.ProductID,
		ProductPrice:       input.ProductPrice,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.ProductPricesUpdateDto) (*dbModel.ProductPrices, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ProductPricesUpdateDto->ProductPrices: input dto cannot be nil")
	}
	return &dbModel.ProductPrices{
		ProductPricePeriod: input.ProductPricePeriod,
		ProductID:          input.ProductID,
		ProductPrice:       input.ProductPrice,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.ProductPrices) ([]dto.ProductPricesDto, error) {
	var outputDtos []dto.ProductPricesDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.ProductPrices) (dto.ProductPricesDto, error) {
	if inputEntity == nil {
		return dto.ProductPricesDto{}, fmt.Errorf("unexpected nil input for mapping between ProductPrices->ProductPricesDto")
	}
	return dto.ProductPricesDto{
		RecordID:           inputEntity.RecordID,
		ProductPricePeriod: inputEntity.ProductPricePeriod,
		ProductID:          inputEntity.ProductID,
		ProductPrice:       inputEntity.ProductPrice,
		CreatedAt:          inputEntity.CreatedAt,
		UpdatedAt:          inputEntity.UpdatedAt,
		Guid:               inputEntity.Guid,
	}, nil
}
