package entities

import "fmt"

type Product struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, Price: %f, Quantity: %d, Status: %t", product.Id, product.Name, product.Price, product.Quantity, product.Status)
}
