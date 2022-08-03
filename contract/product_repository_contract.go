package contract

import (
	"context"
	"database/sql"

	"github.com/mproyyan/go-simple-restful/model"
)

type ProductRepositoryContract interface {
	Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Delete(ctx context.Context, tx *sql.Tx, product model.Product)
	Find(ctx context.Context, tx *sql.Tx, productId int) model.Product
	FindAll(ctx context.Context, tx *sql.Tx) []model.Product
}
