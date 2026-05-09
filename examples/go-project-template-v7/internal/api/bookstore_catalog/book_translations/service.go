package book_translations

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkBookID int64, pkLanguageCode string) (*Dto, error)
	DeleteByID(ctx context.Context, pkBookID int64, pkLanguageCode string) error
	FindByID(ctx context.Context, pkBookID int64, pkLanguageCode string) (*Dto, error)
	FindAll(ctx context.Context) ([]Dto, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Dto, pageable.Page, error)
}

type svc struct {
	repo Repository
}

var _ Service = &svc{}

func NewService(_ context.Context, repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) Save(ctx context.Context, input *CreateDto) (*Dto, error) {
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkBookID int64, pkLanguageCode string) (*Dto, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.repo.UpdateByID(ctx, entityToUpdate, pkBookID, pkLanguageCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *svc) DeleteByID(ctx context.Context, pkBookID int64, pkLanguageCode string) error {
	return s.repo.DeleteByID(ctx, pkBookID, pkLanguageCode)
}

func (s *svc) FindByID(ctx context.Context, pkBookID int64, pkLanguageCode string) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkBookID, pkLanguageCode)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *svc) FindAll(ctx context.Context) ([]Dto, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *svc) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Dto, pageable.Page, error) {
	entities, page, err := s.repo.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *CreateDto) (*BookTranslations, error) {
	if input == nil {
		return nil, fmt.Errorf("convert CreateDto->BookTranslations: input dto cannot be nil")
	}
	return &BookTranslations{
		BookID:          input.BookID,
		LanguageCode:    input.LanguageCode,
		TranslatedTitle: input.TranslatedTitle,
		TranslatedBy:    input.TranslatedBy,
		ReleasedOn:      input.ReleasedOn,
	}, nil
}

func fromUpdateDtoToEntity(input *UpdateDto) (*BookTranslations, error) {
	if input == nil {
		return nil, fmt.Errorf("convert UpdateDto->BookTranslations: input dto cannot be nil")
	}
	return &BookTranslations{
		TranslatedTitle: input.TranslatedTitle,
		TranslatedBy:    input.TranslatedBy,
		ReleasedOn:      input.ReleasedOn,
	}, nil
}

func fromEntitiesToDtos(inputEntities []BookTranslations) ([]Dto, error) {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities { // Iterate using index to avoid copying (gocritic:rangeValCopy)
		toDto, err := fromEntityToDto(&inputEntities[i])
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *BookTranslations) (Dto, error) {
	if inputEntity == nil {
		return Dto{}, fmt.Errorf("unexpected nil input for mapping between BookTranslations->Dto")
	}
	return Dto{
		BookID:          inputEntity.BookID,
		LanguageCode:    inputEntity.LanguageCode,
		TranslatedTitle: inputEntity.TranslatedTitle,
		TranslatedBy:    inputEntity.TranslatedBy,
		ReleasedOn:      inputEntity.ReleasedOn,
	}, nil
}
