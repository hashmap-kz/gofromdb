package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/clients/dto"
	dbModel "go-project-template-v5/internal/api/clients/entity/postgres"

	"go-project-template-v5/internal/api/clients/repository"
	"go-project-template-v5/internal/api/clients/service"
)

type clientsService struct {
	clientsRepository repository.ClientsRepository
}

var _ service.ClientsService = &clientsService{}

func NewClientsService(_ context.Context, clientsRepository repository.ClientsRepository) service.ClientsService {
	return &clientsService{
		clientsRepository: clientsRepository,
	}
}

func (s *clientsService) Save(ctx context.Context, input *dto.ClientsCreateDto) (*dto.ClientsDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.clientsRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *clientsService) UpdateByID(ctx context.Context, input *dto.ClientsUpdateDto, pkRecordID int) (*dto.ClientsDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.clientsRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *clientsService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.clientsRepository.DeleteByID(ctx, pkRecordID)
}

func (s *clientsService) FindByID(ctx context.Context, pkRecordID int) (*dto.ClientsDto, error) {
	entityByID, err := s.clientsRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *clientsService) FindAll(ctx context.Context) ([]dto.ClientsDto, error) {
	entities, err := s.clientsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *clientsService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.ClientsDto, pageable.Page, error) {
	entities, page, err := s.clientsRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.ClientsCreateDto) (*dbModel.Clients, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ClientsCreateDto->Clients: input dto cannot be nil")
	}
	return &dbModel.Clients{
		Email: input.Email,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.ClientsUpdateDto) (*dbModel.Clients, error) {
	if input == nil {
		return nil, fmt.Errorf("convert ClientsUpdateDto->Clients: input dto cannot be nil")
	}
	return &dbModel.Clients{
		Email: input.Email,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Clients) ([]dto.ClientsDto, error) {
	outputDtos := make([]dto.ClientsDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Clients) (dto.ClientsDto, error) {
	if inputEntity == nil {
		return dto.ClientsDto{}, fmt.Errorf("unexpected nil input for mapping between Clients->ClientsDto")
	}
	return dto.ClientsDto{
		RecordID:  inputEntity.RecordID,
		Email:     inputEntity.Email,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		GUID:      inputEntity.GUID,
	}, nil
}
