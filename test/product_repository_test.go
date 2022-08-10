package test

import (
	"context"
	"database/sql"
	"testing"

	mks "github.com/mproyyan/go-simple-restful/mocks"
	"github.com/mproyyan/go-simple-restful/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRepoSaveSuccess(t *testing.T) {
	expectedProduct := model.Product{Id: 1, Name: "Test"}

	mockery := mks.NewProductRepositoryContract(t)
	mockery.On("Save", mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
			return product
		})

	product := mockery.Save(context.Background(), new(sql.Tx), model.Product{Id: 1, Name: "Test"})

	assert.Equal(t, expectedProduct.Id, product.Id)
	assert.Equal(t, expectedProduct.Name, product.Name)
}

func TestRepoUpdateSuccess(t *testing.T) {
	expectedProduct := model.Product{Id: 1, Name: "Updated"}

	mockery := mks.NewProductRepositoryContract(t)
	mockery.On("Update", mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
			return product
		})

	product := mockery.Update(context.Background(), new(sql.Tx), model.Product{Id: 1, Name: "Updated"})

	assert.Equal(t, expectedProduct.Id, product.Id)
	assert.Equal(t, expectedProduct.Name, product.Name)
}

func TestRepoDeleteSuccess(t *testing.T) {
	mockery := mks.NewProductRepositoryContract(t)
	mockery.On("Delete", mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, tx *sql.Tx, product model.Product) bool {
			return true
		})

	result := mockery.Delete(context.Background(), new(sql.Tx), model.Product{Id: 1, Name: "Deleted"})

	assert.True(t, result)
}

func TestRepoFindSuccess(t *testing.T) {
	expectedProduct := model.Product{Id: 7, Name: "Find"}

	mockery := mks.NewProductRepositoryContract(t)
	mockery.On("Find", mock.Anything, mock.Anything, mock.AnythingOfType("int")).
		Return(expectedProduct, nil)

	result, err := mockery.Find(context.Background(), new(sql.Tx), 7)

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

	mockery := mks.NewProductRepositoryContract(t)
	mockery.On("FindAll", mock.Anything, mock.Anything).
		Return(expectedProduct)

	products := mockery.FindAll(context.Background(), new(sql.Tx))
	product := products[0]

	assert.Equal(t, 7, product.Id)
	assert.Equal(t, "FindAll", product.Name)
}
