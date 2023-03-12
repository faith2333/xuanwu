package cicd

type Repo struct {
	TemplateRepo TemplateRepo
}

func NewRepo(templateRepo TemplateRepo) *Repo {
	return &Repo{
		TemplateRepo: templateRepo,
	}
}
