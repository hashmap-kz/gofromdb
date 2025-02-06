package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/steps/dto"
	dbModel "go-project-template-v5/internal/api/steps/entity/postgres"

	"go-project-template-v5/internal/api/steps/repository"
	"go-project-template-v5/internal/api/steps/service"
)

type stepsService struct {
	stepsRepository repository.StepsRepository
}

var _ service.StepsService = &stepsService{}

func NewStepsService(_ context.Context, stepsRepository repository.StepsRepository) service.StepsService {
	return &stepsService{
		stepsRepository: stepsRepository,
	}
}

func (s *stepsService) Save(ctx context.Context, input *dto.StepsCreateDto) (*dto.StepsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.stepsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *stepsService) UpdateByID(ctx context.Context, input *dto.StepsUpdateDto, pkRecordID int) (*dto.StepsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.stepsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *stepsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.stepsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *stepsService) FindByID(ctx context.Context, pkRecordID int) (*dto.StepsDto, error) {
	entityById, err := s.stepsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *stepsService) FindAll(ctx context.Context) ([]dto.StepsDto, error) {
	entities, err := s.stepsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *stepsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.StepsDto, pageable.Page, error) {
	entities, page, err := s.stepsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.StepsCreateDto) (*dbModel.Steps, error) {
	if input == nil {
		return nil, fmt.Errorf("convert StepsCreateDto->Steps: input dto cannot be nil")
	}
	return &dbModel.Steps{
		StepName: input.StepName,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.StepsUpdateDto) (*dbModel.Steps, error) {
	if input == nil {
		return nil, fmt.Errorf("convert StepsUpdateDto->Steps: input dto cannot be nil")
	}
	return &dbModel.Steps{
		StepName: input.StepName,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Steps) ([]dto.StepsDto, error) {
	var outputDtos []dto.StepsDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Steps) (dto.StepsDto, error) {
	if inputEntity == nil {
		return dto.StepsDto{}, fmt.Errorf("unexpected nil input for mapping between Steps->StepsDto")
	}
	return dto.StepsDto{
		RecordID:  inputEntity.RecordID,
		StepName:  inputEntity.StepName,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}
