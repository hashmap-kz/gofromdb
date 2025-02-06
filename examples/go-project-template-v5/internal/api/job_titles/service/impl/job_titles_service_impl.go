package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/job_titles/dto"
	dbModel "go-project-template-v5/internal/api/job_titles/entity/postgres"

	"go-project-template-v5/internal/api/job_titles/repository"
	"go-project-template-v5/internal/api/job_titles/service"
)

type jobTitlesService struct {
	jobTitlesRepository repository.JobTitlesRepository
}

var _ service.JobTitlesService = &jobTitlesService{}

func NewJobTitlesService(_ context.Context, jobTitlesRepository repository.JobTitlesRepository) service.JobTitlesService {
	return &jobTitlesService{
		jobTitlesRepository: jobTitlesRepository,
	}
}

func (s *jobTitlesService) Save(ctx context.Context, input *dto.JobTitlesCreateDto) (*dto.JobTitlesDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.jobTitlesRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *jobTitlesService) UpdateByID(ctx context.Context, input *dto.JobTitlesUpdateDto, pkRecordID int) (*dto.JobTitlesDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.jobTitlesRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *jobTitlesService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.jobTitlesRepository.DeleteByID(ctx, pkRecordID)
}

func (s *jobTitlesService) FindByID(ctx context.Context, pkRecordID int) (*dto.JobTitlesDto, error) {
	entityById, err := s.jobTitlesRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *jobTitlesService) FindAll(ctx context.Context) ([]dto.JobTitlesDto, error) {
	entities, err := s.jobTitlesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *jobTitlesService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.JobTitlesDto, pageable.Page, error) {
	entities, page, err := s.jobTitlesRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.JobTitlesCreateDto) (*dbModel.JobTitles, error) {
	if input == nil {
		return nil, fmt.Errorf("convert JobTitlesCreateDto->JobTitles: input dto cannot be nil")
	}
	return &dbModel.JobTitles{
		TitleName: input.TitleName,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.JobTitlesUpdateDto) (*dbModel.JobTitles, error) {
	if input == nil {
		return nil, fmt.Errorf("convert JobTitlesUpdateDto->JobTitles: input dto cannot be nil")
	}
	return &dbModel.JobTitles{
		TitleName: input.TitleName,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.JobTitles) ([]dto.JobTitlesDto, error) {
	var outputDtos []dto.JobTitlesDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.JobTitles) (dto.JobTitlesDto, error) {
	if inputEntity == nil {
		return dto.JobTitlesDto{}, fmt.Errorf("unexpected nil input for mapping between JobTitles->JobTitlesDto")
	}
	return dto.JobTitlesDto{
		RecordID:  inputEntity.RecordID,
		TitleName: inputEntity.TitleName,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}
