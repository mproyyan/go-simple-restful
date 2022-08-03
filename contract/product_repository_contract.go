package contract

import (
	"context"
	"database/sql"

	"github.com/mproyyan/go-simple-restful/model"
)

type ProductRepositoryContract interface {
	Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Delete(ctx context.Context, tx *sql.Tx, product model.Product) bool
	Find(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Product
}
