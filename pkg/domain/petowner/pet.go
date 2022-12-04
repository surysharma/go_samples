package petowner

import "github.com/learning-go-book/circular_dependency_example/pkg/storage"

type Pet struct {
	Name      string
	Type      string
	OwnerName string
	repo      storage.Repo
}

func (p Pet) Owner() string {
	return p.repo.GetItem(p.Name)
}
