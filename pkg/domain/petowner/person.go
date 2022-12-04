package petowner

import (
	"github.com/circular_dependency_example/pkg/storage"
)

type Person struct {
	Name string
	Repo storage.Repo
	//Uncommenting the below line will cause import cycle issue
	//Logger infra.SimpleLogger
}

func (p Person) Pet() string {
	//Uncommenting the below line will cause import cycle issue
	//p.Logger.AddLog(p.Name)
	return p.Repo.GetItem(p.Name)
}
