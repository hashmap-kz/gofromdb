package tmplts

var ServiceDtosTemplate = `
package dto
import "time"

type {{.DtoName}} struct {
{{- range .DtoFieldsFull}}
	{{.FieldName}} {{.FieldType}}
{{- end}}
}

type {{.DtoCreateName}} struct {
{{- range .DtoFieldsNoPkeysNoDefaults}}
	{{.FieldName}} {{.FieldType}}
{{- end}}
}

type {{.DtoUpdateName}} struct {
{{- range .DtoFieldsNoPkeysNoDefaults}}
	{{.FieldName}} {{.FieldType}}
{{- end}}
}
`

var ServiceInterfaceTemplate = `
package service

import (
	"context"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/{{.PackageName}}/dto"
)

type {{.InterfaceName}} interface {
	Save(ctx context.Context, input *dto.{{.DtoCreateName}}) (*dto.{{.DtoName}}, error)
	GetAll(ctx context.Context) ([]dto.{{.DtoName}}, error)
	GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.{{.DtoName}}, pageable.Page, error)
	Update(ctx context.Context, entityId int, input *dto.{{.DtoUpdateName}}) (*dto.{{.DtoName}}, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*dto.{{.DtoName}}, error)
}
`

var ServiceImplTemplate = `
package impl

import (
	"context"
	"fmt"

	"go-project-template-v5/pkg/pageable"

	"go-project-template-v5/internal/api/{{.PackageName}}/dto"
	dbModel "go-project-template-v5/internal/api/{{.PackageName}}/entity/postgres"

	"go-project-template-v5/internal/api/{{.PackageName}}/repository"
	"go-project-template-v5/internal/api/{{.PackageName}}/service"
)

type {{.ImplName}} struct {
	repo repository.{{.RepositoryName}}
}

var _ service.{{.InterfaceName}} = &{{.ImplName}}{}

func New{{.InterfaceName}}(_ context.Context, repo repository.{{.RepositoryName}}) service.{{.InterfaceName}} {
	return &{{.ImplName}}{repo: repo}
}

func (s *{{.ImplName}}) Save(ctx context.Context, input *dto.{{.DtoCreateName}}) (*dto.{{.DtoName}}, error) {
	save, err := s.repo.Save(ctx, &dbModel.{{.StructName}}{
{{- range .StructFieldsNoPKeys}}
		{{.}}: input.{{.}},
{{- end }}
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

func (s *{{.ImplName}}) GetAll(ctx context.Context) ([]dto.{{.DtoName}}, error) {
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

func (s *{{.ImplName}}) GetAllPaginated(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.{{.DtoName}}, pageable.Page, error) {
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

func (s *{{.ImplName}}) Update(ctx context.Context, entityId int, input *dto.{{.DtoUpdateName}}) (*dto.{{.DtoName}}, error) {
	// update dbModel
	updatedResult, err := s.repo.Update(ctx, entityId, &dbModel.{{.StructName}}{
{{- range .StructFieldsNoPKeys}}
		{{.}}: input.{{.}},
{{- end }}
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

func (s *{{.ImplName}}) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *{{.ImplName}}) GetByID(ctx context.Context, id int) (*dto.{{.DtoName}}, error) {
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

func mapEntitiesToDtos(inputEntities []dbModel.{{.StructName}}) ([]dto.{{.DtoName}}, error) {
	var outputDtos []dto.{{.DtoName}}
	for _, inputEntity := range inputEntities {
		toDto, err := mapEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func mapEntityToDto(inputEntity *dbModel.{{.StructName}}) (dto.{{.DtoName}}, error) {
	if inputEntity == nil {
		return dto.{{.DtoName}}{}, fmt.Errorf("unexpected nil input for mapping between {{.StructName}}->{{.DtoName}}")
	}
	return dto.{{.DtoName}}{
{{- range .StructFieldsWithPKeys}}
		{{.}}: inputEntity.{{.}},
{{- end }}
	}, nil
}
`
