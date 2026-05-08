package app

type GenRepo struct {
	RepoEntity    string
	RepoInterface string
	RepoImpl      string
}

func GenRepository(s TableToStructInfo) GenRepo {
	data := NewRepoTemplateData(s)

	return GenRepo{
		RepoEntity:    PrintFormatted(ExecTemplate("repo_entity", data, FuncMap)),
		RepoInterface: PrintFormatted(ExecTemplate("repo_interface", data, FuncMap)),
		RepoImpl:      PrintFormatted(ExecTemplate("repo_impl", data, FuncMap)),
	}
}
