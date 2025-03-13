package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/users/dto"
	dbModel "go-project-template-v5/internal/api/users/entity/postgres"

	"go-project-template-v5/internal/api/users/repository"
	"go-project-template-v5/internal/api/users/service"
)

type usersService struct {
	usersRepository repository.UsersRepository
}

var _ service.UsersService = &usersService{}

func NewUsersService(_ context.Context, usersRepository repository.UsersRepository) service.UsersService {
	return &usersService{
		usersRepository: usersRepository,
	}
}

func (s *usersService) Save(ctx context.Context, input *dto.UsersCreateDto) (*dto.UsersDto, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	save, err := s.usersRepository.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *usersService) UpdateByID(ctx context.Context, input *dto.UsersUpdateDto, pkRecordID int) (*dto.UsersDto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.usersRepository.UpdateByID(ctx, entityToUpdate, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *usersService) DeleteByID(ctx context.Context, pkRecordID int) error {
	return s.usersRepository.DeleteByID(ctx, pkRecordID)
}

func (s *usersService) FindByID(ctx context.Context, pkRecordID int) (*dto.UsersDto, error) {
	entityByID, err := s.usersRepository.FindByID(ctx, pkRecordID)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *usersService) FindAll(ctx context.Context) ([]dto.UsersDto, error) {
	entities, err := s.usersRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *usersService) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.UsersDto, pageable.Page, error) {
	entities, page, err := s.usersRepository.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.UsersCreateDto) (*dbModel.Users, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UsersCreateDto->Users: input dto cannot be nil")
	}
	return &dbModel.Users{
		Email: input.Email,
	}, nil
}

func fromUpdateDtoToEntity(input *dto.UsersUpdateDto) (*dbModel.Users, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UsersUpdateDto->Users: input dto cannot be nil")
	}
	return &dbModel.Users{
		Email: input.Email,
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.Users) ([]dto.UsersDto, error) {
	outputDtos := make([]dto.UsersDto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.Users) (dto.UsersDto, error) {
	if inputEntity == nil {
		return dto.UsersDto{}, fmt.Errorf("unexpected nil input for mapping between Users->UsersDto")
	}
	return dto.UsersDto{
		RecordID:  inputEntity.RecordID,
		Email:     inputEntity.Email,
		CreatedAt: inputEntity.CreatedAt,
		UpdatedAt: inputEntity.UpdatedAt,
		GUID:      inputEntity.GUID,
	}, nil
}
