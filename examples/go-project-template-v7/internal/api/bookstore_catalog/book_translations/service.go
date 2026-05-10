package book_translations

import (
	"context"
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
	save, err := s.repo.Save(ctx, fromCreateDtoToEntity(input))
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(save)
	return &result, nil
}

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkBookID int64, pkLanguageCode string) (*Dto, error) {
	updated, err := s.repo.UpdateByID(ctx, input, pkBookID, pkLanguageCode)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updated)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkBookID int64, pkLanguageCode string) error {
	return s.repo.DeleteByID(ctx, pkBookID, pkLanguageCode)
}

func (s *svc) FindByID(ctx context.Context, pkBookID int64, pkLanguageCode string) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkBookID, pkLanguageCode)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(entityByID)
	return &result, nil
}

func (s *svc) FindAll(ctx context.Context) ([]Dto, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return fromEntitiesToDtos(entities), nil
}

func (s *svc) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Dto, pageable.Page, error) {
	entities, page, err := s.repo.FindAllPageable(ctx, pq)
	if err != nil {
		return nil, pageable.Page{}, err
	}
	return fromEntitiesToDtos(entities), page, nil
}

// mappers

func fromCreateDtoToEntity(input *CreateDto) *BookTranslations {
	return &BookTranslations{
		BookID:          input.BookID,
		LanguageCode:    input.LanguageCode,
		TranslatedTitle: input.TranslatedTitle,
		TranslatedBy:    input.TranslatedBy,
		ReleasedOn:      input.ReleasedOn,
	}
}

func fromEntitiesToDtos(inputEntities []BookTranslations) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *BookTranslations) Dto {
	return Dto{
		BookID:          inputEntity.BookID,
		LanguageCode:    inputEntity.LanguageCode,
		TranslatedTitle: inputEntity.TranslatedTitle,
		TranslatedBy:    inputEntity.TranslatedBy,
		ReleasedOn:      inputEntity.ReleasedOn,
	}
}
