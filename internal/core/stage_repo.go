package core

import "fmt"

type GenRepo struct {
	Entity     string
	Repository string
}

func GenRepository(s TableToStructInfo) (GenRepo, error) {
	data, err := NewRepoTemplateData(s)
	if err != nil {
		return GenRepo{}, err
	}

	exec := func(name string) (string, error) {
		out, err := ExecTemplate(name, data, FuncMap)
		if err != nil {
			return "", fmt.Errorf("repo %s: %w", name, err)
		}
		return PrintFormatted(out), nil
	}

	entity, err := exec("entity")
	if err != nil {
		return GenRepo{}, err
	}
	repo, err := exec("repository")
	if err != nil {
		return GenRepo{}, err
	}

	return GenRepo{
		Entity:     entity,
		Repository: repo,
	}, nil
}
