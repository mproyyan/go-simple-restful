package contract

import (
	"context"

	"github.com/mproyyan/go-simple-restful/http/request"
	"github.com/mproyyan/go-simple-restful/http/response"
)

type ProductServiceContract interface {
	Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	Delete(ctx context.Context, productId int) bool
	Find(ctx context.Context, productId int) response.ProductResponse
	FindAll(ctx context.Context) []response.ProductResponse
}
