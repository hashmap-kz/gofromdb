package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/client/dto"
	dbModel "go-project-template-v5/internal/api/client/entity/postgres"

	"go-project-template-v5/internal/api/client/repository"
	"go-project-template-v5/internal/api/client/service"
)

type clientService struct {
	repo repository.ClientRepository
}

var _ service.ClientService = &clientService{}

func NewClientService(_ context.Context, repo repository.ClientRepository) service.ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) Save(ctx context.Context, input *dto.ClientCreateDto) (*dto.ClientDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.repo.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *clientService) GetAll(ctx context.Context) ([]dto.ClientDto, error) {
	entities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *clientService) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ClientDto, pageable.Page, error) {
	entities, page, err := s.repo.GetAllPaginated(ctx, pq)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return toDtos, page, nil
}

func (s *clientService) Update(ctx context.Context, entityId int, input *dto.ClientUpdateDto) (*dto.ClientDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.repo.Update(ctx, entityId, entityToUpdate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *clientService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *clientService) GetByID(ctx context.Context, id int) (*dto.ClientDto, error) {
	entityById, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityById)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

// mappers

func fromCreateDtoToEntity(input *dto.ClientCreateDto) (*dbModel.Client, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ClientCreateDto->Client: input dto cannot be nil")
	}
	return &dbModel.Client{
		Email: input.Email,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.ClientUpdateDto) (*dbModel.Client, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ClientUpdateDto->Client: input dto cannot be nil")
	}
	return &dbModel.Client{
		Email: input.Email,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Client) ([]dto.ClientDto, error) {
	var outputDtos []dto.ClientDto
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Client) (dto.ClientDto, error) {
	if inputEntity == nil {
		return dto.ClientDto{}, fmt.Errorf("unexpected nil input for mapping between Client->ClientDto")
	}
	return dto.ClientDto{
		RecordID:  inputEntity.RecordID,
		Email:     inputEntity.Email,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		Guid:      inputEntity.Guid,
	}, nil
}
