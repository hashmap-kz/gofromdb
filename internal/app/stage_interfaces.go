package app

import (
	"genpg-v5/internal/tmplts"
)

type GenInterf struct {
	ServiceInterface string
	RepoInterface    string
	HandlerInterface string
}

func GenInterfaces(structs []TableToStructInfo) GenInterf {
	serviceInterfaceResult := ExecTemplate("service-interface-general", tmplts.ServiceInterfaceGeneral,
		map[string]any{
			"Structs": structs,
		}, FuncMap)

	repoInterfaceResult := ExecTemplate("repo-interface-general", tmplts.RepoInterfaceGeneral,
		map[string]any{
			"Structs": structs,
		}, FuncMap)

	handlerInterfaceResult := ExecTemplate("habdler-interface-general", tmplts.HandlerInterfaceGeneral,
		map[string]any{
			"Structs": structs,
		}, FuncMap)

	return GenInterf{
		ServiceInterface: PrintFormatted(serviceInterfaceResult),
		RepoInterface:    PrintFormatted(repoInterfaceResult),
		HandlerInterface: PrintFormatted(handlerInterfaceResult),
	}
}
