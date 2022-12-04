package infra

import (
	"fmt"
	"github.com/learning-go-book/circular_dependency_example/pkg/domain/petowner"
	"github.com/learning-go-book/circular_dependency_example/pkg/storage"
)

func Handle() {
	logger := SimpleLogger{}
	bob := petowner.Person{
		Name: "Fluffy",
		Repo: storage.NoSQL{},
		//Uncommenting the below line will cause import cycle issue
		//Logger: SimpleLogger{},
	}
	fmt.Println(bob.Pet())
	logger.AddLog(bob.Pet())
}
