package main

import (
	"fmt"
	"inventory/config"
	"inventory/entities"
	"inventory/models"
	"math/rand"
	"time"
)

func main() {
	Demo4()
	Demo5()
	Demo6()
	Demo1()
	Demo2()
	Demo3()
	Demo7()
}

func Demo1() {
	fmt.Println("Demo 1: Find all products")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("=====================================")
			}
		}
	}
}

func Demo2() {
	fmt.Println("Demo 2: Find Products by Name 'Product B'")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		products, err2 := productModel.Search("Product B")
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("=====================================")
			}
		}
	}
}

func Demo3() {
	fmt.Println("Demo 3: Find Products Filtered by Prices '500 - 1500'")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		products, err2 := productModel.FilterByPrices(500, 1500)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("=====================================")
			}
		}
	}
}

func Demo4() {
	fmt.Println("Demo 4: Add Product")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		product := entities.Product{
			Name:     "Product C",
			Price:    1500,
			Quantity: 10,
			Status:   true,
		}

		rowsAffected, err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			fmt.Println("Rows Affected: ", rowsAffected)
			fmt.Println("Product Info")
			product.Id = 0
			fmt.Println(product.ToString())
			fmt.Println("=====================================")
		}
	}
}

func Demo5() {
	fmt.Println("Demo 5: Add Product Method B")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		product := entities.Product{
			Name:     "Product D",
			Price:    2500,
			Quantity: 8,
			Status:   false,
		}

		rowsAffected, err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			fmt.Println("Rows Affected: ", rowsAffected)
			fmt.Println("Product Info")
			product.Id = 0
			fmt.Println(product.ToString())
			fmt.Println("=====================================")
		}
	}
}

func Demo6() {
	fmt.Println("Demo 6: Edit Product by Id '3'")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		// Seed the random number generator with the current time
		rand.Seed(time.Now().UnixNano())

		// Generate a random integer between 97 and 122 (inclusive)
		// These are the ASCII codes for lowercase a-z
		r := rand.Intn(26) + 97

		product := entities.Product{
			Name:     "Product " + string(rune(r)),
			Price:    2500,
			Quantity: 8,
			Status:   false,
			Id:       3,
		}

		rowsAffected, err2 := productModel.Update(&product)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			fmt.Println("Rows Affected: ", rowsAffected)
			fmt.Println("Product Info")
			fmt.Println(product.ToString())
			fmt.Println("=====================================")
		}
	}
}

func Demo7() {
	fmt.Println("Demo 7: Delete Product by Id '3'")
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		rowsAffected, err2 := productModel.Delete(3)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			fmt.Println(err2)
		} else {
			fmt.Println("Rows Affected: ", rowsAffected)
			fmt.Println("Product 3 Deleted")
			fmt.Println("=====================================")
		}
	}
}
