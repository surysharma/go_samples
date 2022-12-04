package petowner

import "github.com/learning-go-book/circular_dependency_example/pkg/storage"

type Person struct {
	Name string
	Repo storage.Repo
}

func (p Person) Pet() string {
	return p.Repo.GetItem(p.Name)
}
