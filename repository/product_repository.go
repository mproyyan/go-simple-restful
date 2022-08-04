package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mproyyan/go-simple-restful/contract"
	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/model"
)

type ProductRepository struct {
}

func NewProductRepository() contract.ProductRepositoryContract {
	return &ProductRepository{}
}

func (pr *ProductRepository) Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	script := "INSERT INTO products(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, script, product.Name)
	helper.CheckErr(err)

	id, _ := result.LastInsertId()
	product.Id = int(id)

	return product
}

func (pr *ProductRepository) Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	script := "UPDATE products SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, product.Name, product.Id)
	helper.CheckErr(err)

	return product
}

func (pr *ProductRepository) Delete(ctx context.Context, tx *sql.Tx, product model.Product) bool {
	script := "DELETE FROM products WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, product.Id)

	if err != nil {
		return false
	}

	return true
}

func (pr *ProductRepository) Find(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error) {
	script := "SELECT id, name FROM products WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, script, productId)
	helper.CheckErr(err)
	defer rows.Close()

	product := model.Product{}
	if rows.Next() {
		scanErr := rows.Scan(&product.Id, &product.Name)
		helper.CheckErr(scanErr)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (pr *ProductRepository) FindAll(ctx context.Context, tx *sql.Tx) []model.Product {
	script := "SELECT id, name FROM products"
	rows, err := tx.QueryContext(ctx, script)
	helper.CheckErr(err)
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Id, &product.Name)
		helper.CheckErr(err)

		products = append(products, product)
	}

	return products
}
