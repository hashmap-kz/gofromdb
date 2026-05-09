package app

type GenInterf struct {
	RepoInterface    string
	ServiceInterface string
	HandlerInterface string
}

func GenInterfaces(structs []TableToStructInfo) GenInterf {
	data := map[string]any{
		"Structs": structs,
	}

	return GenInterf{
		RepoInterface:    PrintFormatted(ExecTemplate("interfaces_repo", data, FuncMap)),
		ServiceInterface: PrintFormatted(ExecTemplate("interfaces_service", data, FuncMap)),
		HandlerInterface: PrintFormatted(ExecTemplate("interfaces_handler", data, FuncMap)),
	}
}
