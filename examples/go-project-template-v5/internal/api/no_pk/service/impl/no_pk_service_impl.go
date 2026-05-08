package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/no_pk/dto"
	dbModel "go-project-template-v5/internal/api/no_pk/entity/postgres"

	"go-project-template-v5/internal/api/no_pk/repository"
	"go-project-template-v5/internal/api/no_pk/service"
)

type noPkService struct {
	noPkRepository repository.NoPkRepository
}

var _ service.NoPkService = &noPkService{}

func NewNoPkService(_ context.Context, noPkRepository repository.NoPkRepository) service.NoPkService {
	return &noPkService{
		noPkRepository: noPkRepository,
	}
}

func (s *noPkService) Save(ctx context.Context, input *dto.NoPkCreateDto) (*dto.NoPkDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.noPkRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *noPkService) FindAll(ctx context.Context) ([]dto.NoPkDto, error) {
	entities, err := s.noPkRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *noPkService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.NoPkDto, pageable.Page, error) {
	entities, page, err := s.noPkRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.NoPkCreateDto) (*dbModel.NoPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert NoPkCreateDto->NoPk: input dto cannot be nil")
	}
	return &dbModel.NoPk{
		EventTime: input.EventTime,
		Message:   input.Message,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.NoPkUpdateDto) (*dbModel.NoPk, error) {
	if input == nil {
		return nil, fmt.Errorf("convert NoPkUpdateDto->NoPk: input dto cannot be nil")
	}
	return &dbModel.NoPk{
		EventTime: input.EventTime,
		Message:   input.Message,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.NoPk) ([]dto.NoPkDto, error) {
	outputDtos := make([]dto.NoPkDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.NoPk) (dto.NoPkDto, error) {
	if inputEntity == nil {
		return dto.NoPkDto{}, fmt.Errorf("unexpected nil input for mapping between NoPk->NoPkDto")
	}
	return dto.NoPkDto{
		EventTime: inputEntity.EventTime,
		Message:   inputEntity.Message,
	}, nil
}
