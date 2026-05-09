package app

import "fmt"

type GenInterf struct {
	RepoInterface    string
	ServiceInterface string
	HandlerInterface string
}

func GenInterfaces(structs []TableToStructInfo) (GenInterf, error) {
	data := map[string]any{
		"Structs": structs,
	}

	exec := func(name string) (string, error) {
		s, err := ExecTemplate(name, data, FuncMap)
		if err != nil {
			return "", fmt.Errorf("interface %s: %w", name, err)
		}
		return PrintFormatted(s), nil
	}

	repoIface, err := exec("interfaces_repo")
	if err != nil {
		return GenInterf{}, err
	}
	svcIface, err := exec("interfaces_service")
	if err != nil {
		return GenInterf{}, err
	}
	handlerIface, err := exec("interfaces_handler")
	if err != nil {
		return GenInterf{}, err
	}

	return GenInterf{
		RepoInterface:    repoIface,
		ServiceInterface: svcIface,
		HandlerInterface: handlerIface,
	}, nil
}
