package file

type Repo struct {
	fileName string
}

func New(fileName string) *Repo {
	return &Repo{fileName: fileName}
}