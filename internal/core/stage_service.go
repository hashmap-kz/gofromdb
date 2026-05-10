package core

import "fmt"

type GenSvc struct {
	Dto     string
	Service string
}

func GenService(s TableToStructInfo) (GenSvc, error) {
	data, err := NewServiceTemplateData(s)
	if err != nil {
		return GenSvc{}, err
	}

	exec := func(name string) (string, error) {
		out, err := ExecTemplate(name, data)
		if err != nil {
			return "", fmt.Errorf("service %s: %w", name, err)
		}
		return out, nil
	}

	dto, err := exec("dto")
	if err != nil {
		return GenSvc{}, err
	}
	svc, err := exec("service")
	if err != nil {
		return GenSvc{}, err
	}

	return GenSvc{
		Dto:     dto,
		Service: svc,
	}, nil
}
