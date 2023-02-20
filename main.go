package main

import (
	"fmt"
)

type ColorType int

const (
	BLUE ColorType = iota
	RED
	GREEN
	ORANGE
)

type Product struct {
	Name        string
	Description string
	Colour      ColorType
}

type FilterStrategy interface {
	isPresent(p *Product) bool
}

type ColourFilter struct {
	colour ColorType
}

func (cf *ColourFilter) isPresent(p *Product) bool {
	return cf.colour == p.Colour

}

type NameFilter struct {
	Name string
}

func (cf *NameFilter) isPresent(p *Product) bool {
	return cf.Name == p.Name

}

func filter(filter FilterStrategy, p []Product) {
	for _, v := range p {

		if filter.isPresent(&v) {
			fmt.Println("Product name is", v.Name, ",Description is", v.Description)
		}
	}
}

func main() {

	products := make([]Product, 0, 10)

	products = append(products, Product{Name: "Apple", Colour: RED, Description: "Gala Apples"})
	products = append(products, Product{Name: "Banana", Colour: GREEN, Description: "Green Carribiean Bananas"})
	products = append(products, Product{Name: "Orange", Colour: ORANGE, Description: "Samartian Oranges"})
	products = append(products, Product{Name: "Apple", Colour: RED, Description: "Indian Apples"})
	products = append(products, Product{Name: "Banana", Colour: GREEN, Description: "Indian Bananas"})
	products = append(products, Product{Name: "Orange", Colour: ORANGE, Description: "Jamican Oranges"})

	colorFilter := ColourFilter{GREEN}
	filter(&colorFilter, products)

	nameFilter := NameFilter{Name: "Orange"}
	filter(&nameFilter, products)

}
