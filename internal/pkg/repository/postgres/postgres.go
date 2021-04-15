package postgres

type Repo struct{}

func New() *Repo {
	return &Repo{}
}
