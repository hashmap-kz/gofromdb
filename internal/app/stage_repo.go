package app

type GenRepo struct {
	Entity     string
	Repository string
}

func GenRepository(s TableToStructInfo) GenRepo {
	data := NewRepoTemplateData(s)

	return GenRepo{
		Entity:     PrintFormatted(ExecTemplate("entity", data, FuncMap)),
		Repository: PrintFormatted(ExecTemplate("repository", data, FuncMap)),
	}
}
