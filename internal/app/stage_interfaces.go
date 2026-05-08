package app

type GenInterf struct {
	ServiceInterface string
	RepoInterface    string
	HandlerInterface string
}

func GenInterfaces(structs []TableToStructInfo) GenInterf {
	data := map[string]any{
		"Structs": structs,
	}

	return GenInterf{
		ServiceInterface: PrintFormatted(ExecTemplate("interfaces_service", data, FuncMap)),
		RepoInterface:    PrintFormatted(ExecTemplate("interfaces_repo", data, FuncMap)),
		HandlerInterface: PrintFormatted(ExecTemplate("interfaces_handler", data, FuncMap)),
	}
}
