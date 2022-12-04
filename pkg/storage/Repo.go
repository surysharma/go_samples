package storage

import (
	"log"
)

type Repo interface {
	StoreItem(item string)
	GetItem(item string) string
}

var ownersRepo = map[string]string{
	"Bob":   "Fluffy",
	"Julia": "Rex",
}

type NoSQL struct{}

func (n NoSQL) StoreItem(item string) {
	log.Println("Item added in Repo...", item)
}
func (n NoSQL) GetItem(item string) string {
	return ownersRepo[item]
}
