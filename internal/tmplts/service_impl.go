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
	UpdateByID(ctx context.Context, input *dto.{{.DtoUpdateName}}, {{.ParametersByPkeys}}) (*dto.{{.DtoName}}, error)
	DeleteByID(ctx context.Context, {{.ParametersByPkeys}}) error
	FindByID(ctx context.Context, {{.ParametersByPkeys}}) (*dto.{{.DtoName}}, error)
	FindAll(ctx context.Context) ([]dto.{{.DtoName}}, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.{{.DtoName}}, pageable.Page, error)
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
	{{.RepositoryVarName}} repository.{{.RepositoryInterfaceName}}
}

var _ service.{{.InterfaceName}} = &{{.ImplName}}{}

func New{{.InterfaceName}}(_ context.Context, {{.RepositoryVarName}} repository.{{.RepositoryInterfaceName}}) service.{{.InterfaceName}} {
	return &{{.ImplName}}{
		{{.RepositoryVarName}}: {{.RepositoryVarName}},
	}
}

func (s *{{.ImplName}}) Save(ctx context.Context, input *dto.{{.DtoCreateName}}) (*dto.{{.DtoName}}, error) {
	entityToCreate, err := fromCreateDtoToEntity(input)
	if err != nil {
		return nil, err
	}	
	save, err := s.{{.RepositoryVarName}}.Save(ctx, entityToCreate)
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(save)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *{{.ImplName}}) UpdateByID(ctx context.Context, input *dto.{{.DtoUpdateName}}, {{.ParametersByPkeys}}) (*dto.{{.DtoName}}, error) {
	entityToUpdate, err := fromUpdateDtoToEntity(input)
	if err != nil {
		return nil, err
	}
	updatedResult, err := s.{{.RepositoryVarName}}.UpdateByID(ctx, entityToUpdate, {{.ArgumentsByPkeys}})
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(updatedResult)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *{{.ImplName}}) DeleteByID(ctx context.Context, {{.ParametersByPkeys}}) error {
	return s.{{.RepositoryVarName}}.DeleteByID(ctx, {{.ArgumentsByPkeys}})
}

func (s *{{.ImplName}}) FindByID(ctx context.Context, {{.ParametersByPkeys}}) (*dto.{{.DtoName}}, error) {
	entityByID, err := s.{{.RepositoryVarName}}.FindByID(ctx, {{.ArgumentsByPkeys}})
	if err != nil {
		return nil, err
	}
	toDto, err := fromEntityToDto(entityByID)
	if err != nil {
		return nil, err
	}
	return &toDto, err
}

func (s *{{.ImplName}}) FindAll(ctx context.Context) ([]dto.{{.DtoName}}, error) {
	entities, err := s.{{.RepositoryVarName}}.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	toDtos, err := fromEntitiesToDtos(entities)
	if err != nil {
		return nil, err
	}
	return toDtos, nil
}

func (s *{{.ImplName}}) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]dto.{{.DtoName}}, pageable.Page, error) {
	entities, page, err := s.{{.RepositoryVarName}}.FindAllPageable(ctx, pq)
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

func fromCreateDtoToEntity(input *dto.{{.DtoCreateName}}) (*dbModel.{{.StructName}}, error) {
	if input == nil {
		return nil, fmt.Errorf("convert {{.DtoCreateName}}->{{.StructName}}: input dto cannot be nil")
	}
	return &dbModel.{{.StructName}}{
{{- range .DtoFieldsNoPkeysNoDefaults}}
		{{.FieldName}}: input.{{.FieldName}},
{{- end}}
	}, nil
}

func fromUpdateDtoToEntity(input *dto.{{.DtoUpdateName}}) (*dbModel.{{.StructName}}, error) {
	if input == nil {
		return nil, fmt.Errorf("convert {{.DtoUpdateName}}->{{.StructName}}: input dto cannot be nil")
	}
	return &dbModel.{{.StructName}}{
{{- range .DtoFieldsNoPkeysNoDefaults}}
		{{.FieldName}}: input.{{.FieldName}},
{{- end}}
	}, nil
}

func fromEntitiesToDtos(inputEntities []dbModel.{{.StructName}}) ([]dto.{{.DtoName}}, error) {
	var outputDtos []dto.{{.DtoName}}
	for _, inputEntity := range inputEntities {
		toDto, err := fromEntityToDto(&inputEntity)
		if err != nil {
			return nil, err
		}
		outputDtos = append(outputDtos, toDto)
	}
	return outputDtos, nil
}

func fromEntityToDto(inputEntity *dbModel.{{.StructName}}) (dto.{{.DtoName}}, error) {
	if inputEntity == nil {
		return dto.{{.DtoName}}{}, fmt.Errorf("unexpected nil input for mapping between {{.StructName}}->{{.DtoName}}")
	}
	return dto.{{.DtoName}}{
{{- range .DtoFieldsFull}}
		{{.FieldName}}: inputEntity.{{.FieldName}},
{{- end}}
	}, nil
}
`
