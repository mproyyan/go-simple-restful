package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mproyyan/go-simple-restful/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	mock.Mock
}

var productRepoMock = ProductRepoMock{mock.Mock{}}

func (pr *ProductRepoMock) Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	arguments := pr.Called(ctx, tx, product)

	if arguments.Get(0) == nil {
		panic("no return value")
	}

	data := arguments.Get(0).(model.Product)
	return data
}

func (pr *ProductRepoMock) Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	arguments := pr.Called(ctx, tx, product)

	if arguments.Get(0) == nil {
		panic("no return value")
	}

	data := arguments.Get(0).(model.Product)
	return data
}

func (pr *ProductRepoMock) Delete(ctx context.Context, tx *sql.Tx, product model.Product) bool {
	arguments := pr.Called(ctx, tx, product)

	if arguments.Get(0) == nil {
		panic("no return value")
	}

	data := arguments.Get(0).(bool)
	return data
}

func (pr *ProductRepoMock) Find(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error) {
	arguments := pr.Called(ctx, tx, productId)

	if arguments.Get(0) == nil {
		panic("no return value")
	}

	data := arguments.Get(0).(model.Product)
	return data, nil
}

func (pr *ProductRepoMock) FindAll(ctx context.Context, tx *sql.Tx) []model.Product {
	arguments := pr.Called(ctx, tx)

	if arguments.Get(0) == nil {
		panic("no return value")
	}

	data := arguments.Get(0).([]model.Product)
	return data
}

func TestRepoSaveSuccess(t *testing.T) {
	expectedProduct := model.Product{Id: 1, Name: "Test"}

	productRepoMock.
		On("Save", context.Background(), new(sql.Tx), model.Product{Name: "Test"}).
		Return(expectedProduct)

	product := productRepoMock.Save(context.Background(), new(sql.Tx), model.Product{Name: "Test"})

	assert.Equal(t, expectedProduct.Id, product.Id)
	assert.Equal(t, expectedProduct.Name, product.Name)
}

func TestRepoUpdateSuccess(t *testing.T) {
	expectedProduct := model.Product{Id: 1, Name: "Updated"}

	productRepoMock.
		On("Update", context.Background(), new(sql.Tx), expectedProduct).
		Return(expectedProduct)

	product := productRepoMock.Update(context.Background(), new(sql.Tx), model.Product{Id: 1, Name: "Updated"})

	assert.Equal(t, expectedProduct.Id, product.Id)
	assert.Equal(t, expectedProduct.Name, product.Name)
}

func TestRepoDeleteSuccess(t *testing.T) {
	productRepoMock.
		On("Delete", context.Background(), new(sql.Tx), model.Product{Id: 1, Name: "Deleted"}).
		Return(true)

	result := productRepoMock.Delete(context.Background(), new(sql.Tx), model.Product{Id: 1, Name: "Deleted"})

	assert.True(t, result)
}

func TestRepoFindSuccess(t *testing.T) {
	expectedProduct := model.Product{Id: 7, Name: "Find"}

	productRepoMock.
		On("Find", context.Background(), new(sql.Tx), 7).
		Return(expectedProduct)

	result, err := productRepoMock.Find(context.Background(), new(sql.Tx), 7)

	assert.Nil(t, err)
	assert.Equal(t, expectedProduct.Id, result.Id)
	assert.Equal(t, expectedProduct.Name, result.Name)
}

func TestRepoFindAllSucces(t *testing.T) {
	expectedProduct := []model.Product{
		{
			Id:   7,
			Name: "FindAll",
		},
	}

	productRepoMock.
		On("FindAll", context.Background(), new(sql.Tx)).
		Return(expectedProduct)

	products := productRepoMock.FindAll(context.Background(), new(sql.Tx))
	product := products[0]

	assert.Equal(t, 7, product.Id)
	assert.Equal(t, "FindAll", product.Name)
}
