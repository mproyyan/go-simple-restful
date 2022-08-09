package test

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/mproyyan/go-simple-restful/http/request"
	"github.com/mproyyan/go-simple-restful/model"
	"github.com/mproyyan/go-simple-restful/repository"
	"github.com/mproyyan/go-simple-restful/service"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func NewService(db *sql.DB) *service.ProductService {
	var service = service.NewProductService(
		repository.NewProductRepository(),
		db,
		validator.New(),
	)

	return service
}

func ExpectationFulfilled(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

var p *model.Product = &model.Product{
	Id:   1,
	Name: "Product Test",
}

func TestCreateProductSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "INSERT INTO products"

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(p.Name).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	service := NewService(db)

	product := service.Create(
		context.Background(),
		request.ProductCreateRequest{Name: "Product Test"},
	)

	assert.Equal(t, 1, product.Id)
	assert.Equal(t, "Product Test", product.Name)

	ExpectationFulfilled(t, mock)
}

func TestCreateProductError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "INSERT INTO products"

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(p.Name).WillReturnError(errors.New("Create Failed"))
	mock.ExpectRollback()

	service := NewService(db)

	service.Create(
		context.Background(),
		request.ProductCreateRequest{Name: "Product Test"},
	)

	ExpectationFulfilled(t, mock)
}

func TestUpdateProductSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	updateQuery := regexp.QuoteMeta(`UPDATE products SET name = ? WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.Id, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.Id).WillReturnRows(data)
	mock.ExpectExec(updateQuery).WithArgs("Updated", 1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	service := NewService(db)

	product := service.Update(
		context.Background(),
		request.ProductUpdateRequest{Id: 1, Name: "Updated"},
	)

	assert.Equal(t, 1, product.Id)
	assert.Equal(t, "Updated", product.Name)

	ExpectationFulfilled(t, mock)
}

func TestUpdateProductError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	updateQuery := regexp.QuoteMeta(`UPDATE products SET name = ? WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.Id, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.Id).WillReturnRows(data)
	mock.ExpectExec(updateQuery).WithArgs("Updated", 1).WillReturnError(errors.New("Update gagal"))
	mock.ExpectRollback()

	service := NewService(db)

	service.Update(
		context.Background(),
		request.ProductUpdateRequest{Id: 1, Name: "Updated"},
	)

	ExpectationFulfilled(t, mock)
}

func TestDeleteProductSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	deleteQuery := regexp.QuoteMeta(`DELETE FROM products WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.Id, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.Id).WillReturnRows(data)
	mock.ExpectExec(deleteQuery).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	service := NewService(db)

	result := service.Delete(
		context.Background(),
		1,
	)

	assert.True(t, result)

	ExpectationFulfilled(t, mock)
}

func TestDeleteProductError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	deleteQuery := regexp.QuoteMeta(`DELETE FROM products WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.Id, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.Id).WillReturnRows(data)
	mock.ExpectExec(deleteQuery).WithArgs(1).WillReturnError(errors.New("Delete failed"))
	mock.ExpectCommit()

	service := NewService(db)

	result := service.Delete(
		context.Background(),
		1,
	)

	assert.False(t, result)

	ExpectationFulfilled(t, mock)
}

func TestFindProductSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.Id, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.Id).WillReturnRows(data)
	mock.ExpectCommit()

	service := NewService(db)

	product := service.Find(
		context.Background(),
		1,
	)

	assert.Equal(t, p.Id, product.Id)
	assert.Equal(t, p.Name, product.Name)

	ExpectationFulfilled(t, mock)
}

func TestFindProductError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.Id).WillReturnError(errors.New("Product not found"))
	mock.ExpectRollback()

	service := NewService(db)

	service.Find(
		context.Background(),
		1,
	)

	ExpectationFulfilled(t, mock)
}

func TestFindAllProductSuccess(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findAllQuery := regexp.QuoteMeta(`SELECT id, name FROM products`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).
		AddRow(p.Id, p.Name).
		AddRow(p.Id, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findAllQuery).WillReturnRows(data)
	mock.ExpectCommit()

	service := NewService(db)
	products := service.FindAll(context.Background())

	assert.Len(t, products, 2)
	ExpectationFulfilled(t, mock)
}

func TestFindAllProductError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	findAllQuery := regexp.QuoteMeta(`SELECT id, name FROM products`)

	mock.ExpectBegin()
	mock.ExpectQuery(findAllQuery).WillReturnError(errors.New("Error coy"))
	mock.ExpectRollback()

	service := NewService(db)
	service.FindAll(context.Background())

	ExpectationFulfilled(t, mock)
}
