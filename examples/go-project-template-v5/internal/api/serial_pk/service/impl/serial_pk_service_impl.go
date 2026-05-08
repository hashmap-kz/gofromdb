package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/serial_pk/dto"
	dbModel "go-project-template-v5/internal/api/serial_pk/entity/postgres"

	"go-project-template-v5/internal/api/serial_pk/repository"
	"go-project-template-v5/internal/api/serial_pk/service"
)

type serialPkService struct {
	serialPkRepository repository.SerialPkRepository
}

var _ service.SerialPkService = &serialPkService{}

func NewSerialPkService(_ context.Context, serialPkRepository repository.SerialPkRepository) service.SerialPkService {
	return &serialPkService{
		serialPkRepository: serialPkRepository,
	}
}

func (s *serialPkService) Save(ctx context.Context, input *dto.SerialPkCreateDto) (*dto.SerialPkDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.serialPkRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *serialPkService) UpdateByID(ctx context.Context, input *dto.SerialPkUpdateDto, pkID int64) (*dto.SerialPkDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.serialPkRepository.UpdateByID(ctx, entityToUpdate, pkID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *serialPkService) DeleteByID(ctx context.Context, pkID int64) error {
	return s.serialPkRepository.DeleteByID(ctx, pkID)
}

func (s *serialPkService) FindByID(ctx context.Context, pkID int64) (*dto.SerialPkDto, error) {
	entityByID, err := s.serialPkRepository.FindByID(ctx, pkID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *serialPkService) FindAll(ctx context.Context) ([]dto.SerialPkDto, error) {
	entities, err := s.serialPkRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *serialPkService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.SerialPkDto, pageable.Page, error) {
	entities, page, err := s.serialPkRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.SerialPkCreateDto) (*dbModel.SerialPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert SerialPkCreateDto->SerialPk: input dto cannot be nil")
	}
	return &dbModel.SerialPk{
		Name: input.Name,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.SerialPkUpdateDto) (*dbModel.SerialPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert SerialPkUpdateDto->SerialPk: input dto cannot be nil")
	}
	return &dbModel.SerialPk{
		Name: input.Name,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.SerialPk) ([]dto.SerialPkDto, error) {
	outputDtos := make([]dto.SerialPkDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.SerialPk) (dto.SerialPkDto, error) {
	if inputEntity == nil {
		return dto.SerialPkDto{}, fmt.Errorf("unexpected nil input for mapping between SerialPk->SerialPkDto")
	}
	return dto.SerialPkDto{
		ID:   inputEntity.ID,
		Name: inputEntity.Name,
	}, nil
}
