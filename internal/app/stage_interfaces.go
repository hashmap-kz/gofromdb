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
	data := map[string]any{
		"Structs": structs,
	}

	serviceInterfaceResult := ExecTemplate("service-interface-general", tmplts.ServiceInterfaceGeneral, data, FuncMap)
	repoInterfaceResult := ExecTemplate("repo-interface-general", tmplts.RepoInterfaceGeneral, data, FuncMap)
	handlerInterfaceResult := ExecTemplate("handler-interface-general", tmplts.HandlerInterfaceGeneral, data, FuncMap)

	return GenInterf{
		ServiceInterface: PrintFormatted(serviceInterfaceResult),
		RepoInterface:    PrintFormatted(repoInterfaceResult),
		HandlerInterface: PrintFormatted(handlerInterfaceResult),
	}
}
