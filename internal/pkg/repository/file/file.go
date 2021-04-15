package file

type Repo struct {
	fileName string
}

func NewRepo(fileName string) *Repo {
	return &Repo{fileName: fileName}
}