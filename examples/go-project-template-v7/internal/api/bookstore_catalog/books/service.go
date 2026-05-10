package books

import (
	"context"
	"go-project-template-v7/pkg/pageable"
)

type Service interface {
	Save(ctx context.Context, input *CreateDto) (*Dto, error)
	UpdateByID(ctx context.Context, input *UpdateDto, pkBookID int64) (*Dto, error)
	DeleteByID(ctx context.Context, pkBookID int64) error
	FindByID(ctx context.Context, pkBookID int64) (*Dto, error)
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

func (s *svc) UpdateByID(ctx context.Context, input *UpdateDto, pkBookID int64) (*Dto, error) {
	updatedResult, err := s.repo.UpdateByID(ctx, fromUpdateDtoToEntity(input), pkBookID)
	if err != nil {
		return nil, err
	}
	result := fromEntityToDto(updatedResult)
	return &result, nil
}

func (s *svc) DeleteByID(ctx context.Context, pkBookID int64) error {
	return s.repo.DeleteByID(ctx, pkBookID)
}

func (s *svc) FindByID(ctx context.Context, pkBookID int64) (*Dto, error) {
	entityByID, err := s.repo.FindByID(ctx, pkBookID)
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

func fromCreateDtoToEntity(input *CreateDto) *Books {
	return &Books{
		PublisherCode: input.PublisherCode,
		Isbn13:        input.Isbn13,
		Title:         input.Title,
		Subtitle:      input.Subtitle,
		Description:   input.Description,
		Price:         input.Price,
		WeightGrams:   input.WeightGrams,
		Rating:        input.Rating,
		PublishedOn:   input.PublishedOn,
		Tags:          input.Tags,
		Attrs:         input.Attrs,
		CoverImage:    input.CoverImage,
		ArchivedAt:    input.ArchivedAt,
	}
}

func fromUpdateDtoToEntity(input *UpdateDto) *Books {
	return &Books{
		PublisherCode: input.PublisherCode,
		Isbn13:        input.Isbn13,
		Title:         input.Title,
		Subtitle:      input.Subtitle,
		Description:   input.Description,
		Price:         input.Price,
		WeightGrams:   input.WeightGrams,
		Rating:        input.Rating,
		PublishedOn:   input.PublishedOn,
		Tags:          input.Tags,
		Attrs:         input.Attrs,
		CoverImage:    input.CoverImage,
		ArchivedAt:    input.ArchivedAt,
	}
}

func fromEntitiesToDtos(inputEntities []Books) []Dto {
	outputDtos := make([]Dto, 0, len(inputEntities))
	for i := range inputEntities {
		outputDtos = append(outputDtos, fromEntityToDto(&inputEntities[i]))
	}
	return outputDtos
}

func fromEntityToDto(inputEntity *Books) Dto {
	return Dto{
		BookID:        inputEntity.BookID,
		PublisherCode: inputEntity.PublisherCode,
		Isbn13:        inputEntity.Isbn13,
		Title:         inputEntity.Title,
		Subtitle:      inputEntity.Subtitle,
		Description:   inputEntity.Description,
		Price:         inputEntity.Price,
		WeightGrams:   inputEntity.WeightGrams,
		Rating:        inputEntity.Rating,
		PublishedOn:   inputEntity.PublishedOn,
		Tags:          inputEntity.Tags,
		Attrs:         inputEntity.Attrs,
		CoverImage:    inputEntity.CoverImage,
		ArchivedAt:    inputEntity.ArchivedAt,
		TitleSearch:   inputEntity.TitleSearch,
	}
}
