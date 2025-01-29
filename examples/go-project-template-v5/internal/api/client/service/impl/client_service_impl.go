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
	save, err := s.repo.Save(ctx, &dbModel.Client{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	toDto, err := mapEntityToDto(save)
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
	toDtos, err := mapEntitiesToDtos(entities)
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
	toDtos, err := mapEntitiesToDtos(entities)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return toDtos, page, nil
}

func (s *clientService) Update(ctx context.Context, input *dto.ClientUpdateDto) (*dto.ClientDto, error) {
	// update dbModel
	updatedResult, err := s.repo.Update(ctx, &dbModel.Client{
		RecordID: input.RecordID,
		Email:    input.Email,
	})
	if err != nil {
		return nil, err
	}

	// convert dbModel to internal dto
	toDto, err := mapEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}

	return &toDto, err
}

func (s *clientService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *clientService) GetByID(ctx context.Context, id int) (*dto.ClientDto, error) {
	// retrieve dbModel by id
	entityById, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// convert dbModel to internal dto
	toDto, err := mapEntityToDto(entityById)
	if err != nil {
		return nil, err
	}

	return &toDto, err
}

// mappers

func mapEntitiesToDtos(inputEntities []dbModel.Client) ([]dto.ClientDto, error) {
	var outputDtos []dto.ClientDto
	for _, inputEntity := range inputEntities {
		toDto, err := mapEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func mapEntityToDto(inputEntity *dbModel.Client) (dto.ClientDto, error) {
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
