package cicd

type Repo struct {
	TemplateRepo TemplateRepo
}

func NewRepo() *Repo {
	return &Repo{}
}
