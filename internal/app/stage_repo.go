package app

import "genpg-v5/internal/tmplts"

type GenRepo struct {
	RepoEntity    string
	RepoInterface string
	RepoImpl      string
}

func GenRepository(s TableToStructInfo) GenRepo {
	data := NewRepoTemplateData(s)

	interfaceRes := ExecTemplate("entity-interface", tmplts.RepoInterfaceTemplate, data, FuncMap)
	modelsRes := ExecTemplate("entity", tmplts.EntityTemplate, data, FuncMap)
	implRes := ExecTemplate("funcs", tmplts.RepoImplTemplate, data, FuncMap)

	return GenRepo{
		RepoEntity:    PrintFormatted(modelsRes),
		RepoInterface: PrintFormatted(interfaceRes),
		RepoImpl:      PrintFormatted(implRes),
	}
}
