package models

import (
	"database/sql"
	"inventory/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("SELECT * FROM [inventory].[dbo].[Products]")

	if err != nil {
		return nil, err
	} else {
		return ProceedProduct(rows)
	}
}

func (productModel ProductModel) Search(keyword string) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("SELECT * FROM [inventory].[dbo].[Products] WHERE name LIKE ?",
		"%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {
		return ProceedProduct(rows)
	}
}

func (productModel ProductModel) FilterByPrices(min, max float64) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("SELECT * FROM [inventory].[dbo].[Products] WHERE price BETWEEN ? AND ?",
		min, max)

	if err != nil {
		return nil, err
	} else {
		return ProceedProduct(rows)
	}
}

func (productModel ProductModel) Create(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("INSERT INTO [inventory].[dbo].[Products] VALUES (?, ?, ?, ?)",
		product.Name, product.Price, product.Quantity, product.Status)

	if err != nil {
		return 0, err
	} else {
		product.Id, _ = result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
}

func (productModel ProductModel) Create2(product *entities.Product) error {
	row, err := productModel.Db.Query("INSERT INTO [inventory].[dbo].[Products] VALUES (?, ?, ?, ?) SELECT CONVERT(bigint, SCOPE_IDENTITY());",
		product.Name, product.Price, product.Quantity, product.Status)

	if err != nil {
		return err
	} else {
		var newId int64
		row.Next()
		row.Scan(&newId)
		product.Id = newId
		return nil
	}
}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("UPDATE [inventory].[dbo].[Products] SET name = ?, price = ?, quantity = ?, status = ? WHERE id = ?",
		product.Name, product.Price, product.Quantity, product.Status, product.Id)

	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("DELETE FROM [inventory].[dbo].[Products] WHERE id = ?", id)

	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
}

func ProceedProduct(rows *sql.Rows) ([]entities.Product, error) {
	var products []entities.Product
	for rows.Next() {
		var id int64
		var name string
		var price float64
		var quantity int
		var status bool

		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
		if err2 != nil {
			return nil, err2
		} else {
			product := entities.Product{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
				Status:   status,
			}

			products = append(products, product)
		}
	}

	return products, nil
}
