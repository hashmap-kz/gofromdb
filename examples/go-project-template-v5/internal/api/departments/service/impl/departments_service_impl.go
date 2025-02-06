package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/departments/dto"
	dbModel "go-project-template-v5/internal/api/departments/entity/postgres"

	"go-project-template-v5/internal/api/departments/repository"
	"go-project-template-v5/internal/api/departments/service"
)

type departmentsService struct {
	departmentsRepository repository.DepartmentsRepository
}

var _ service.DepartmentsService = &departmentsService{}

func NewDepartmentsService(_ context.Context, departmentsRepository repository.DepartmentsRepository) service.DepartmentsService {
	return &departmentsService{
		departmentsRepository: departmentsRepository,
	}
}

func (s *departmentsService) Save(ctx context.Context, input *dto.DepartmentsCreateDto) (*dto.DepartmentsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.departmentsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *departmentsService) UpdateByID(ctx context.Context, input *dto.DepartmentsUpdateDto, pkRecordID int) (*dto.DepartmentsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.departmentsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *departmentsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.departmentsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *departmentsService) FindByID(ctx context.Context, pkRecordID int) (*dto.DepartmentsDto, error) {
	entityById, err := s.departmentsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *departmentsService) FindAll(ctx context.Context) ([]dto.DepartmentsDto, error) {
	entities, err := s.departmentsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *departmentsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.DepartmentsDto, pageable.Page, error) {
	entities, page, err := s.departmentsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.DepartmentsCreateDto) (*dbModel.Departments, error) {
	if input == nil {
		return nil, fmt.Errorf("convert DepartmentsCreateDto->Departments: input dto cannot be nil")
	}
	return &dbModel.Departments{
		DepartmentName: input.DepartmentName,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.DepartmentsUpdateDto) (*dbModel.Departments, error) {
	if input == nil {
		return nil, fmt.Errorf("convert DepartmentsUpdateDto->Departments: input dto cannot be nil")
	}
	return &dbModel.Departments{
		DepartmentName: input.DepartmentName,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Departments) ([]dto.DepartmentsDto, error) {
	var outputDtos []dto.DepartmentsDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Departments) (dto.DepartmentsDto, error) {
	if inputEntity == nil {
		return dto.DepartmentsDto{}, fmt.Errorf("unexpected nil input for mapping between Departments->DepartmentsDto")
	}
	return dto.DepartmentsDto{
		RecordID:       inputEntity.RecordID,
		DepartmentName: inputEntity.DepartmentName,
		CreatedAt:      inputEntity.CreatedAt,
		UpdatedAt:      inputEntity.UpdatedAt,
		Guid:           inputEntity.Guid,
	}, nil
}
