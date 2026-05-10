package core

import "fmt"

type GenHandl struct {
	Payload string
	Handler string
}

func GenHandler(s TableToStructInfo) (GenHandl, error) {
	data, err := NewHandlerTemplateData(s)
	if err != nil {
		return GenHandl{}, err
	}

	exec := func(name string) (string, error) {
		out, err := ExecTemplate(name, data)
		if err != nil {
			return "", fmt.Errorf("handler %s: %w", name, err)
		}
		return out, nil
	}

	payload, err := exec("payload")
	if err != nil {
		return GenHandl{}, err
	}
	handler, err := exec("handler")
	if err != nil {
		return GenHandl{}, err
	}

	return GenHandl{
		Payload: payload,
		Handler: handler,
	}, nil
}
