package http

import (
	"fmt"
	httpresponse "golang-starter/internal/protocols/http/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h HttpHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get all products
	// products := h.ProductService.GetProducts()

	// productsResponse := dto.ParseFromEntities(products)

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", productsResponse)
	httpresponse.Text(w, http.StatusOK, "OK")
}

func (h HttpHandlerImpl) GetProductByID(w http.ResponseWriter, r *http.Request) {
	rawProduct_id := chi.URLParam(r, "product_id")
	product_id, err := strconv.Atoi(rawProduct_id)
	if err != nil {
		httpresponse.Json(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// product := h.ProductService.GetProductByProductID(product_id)

	// productResponse := dto.ParseFromEntity(product)

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", productResponse)
	httpresponse.Text(w, http.StatusOK, fmt.Sprintf("%d", product_id))
}

func (h HttpHandlerImpl) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	// rawProduct_id := ctx.Params("product_id")
	// product_id, err := strconv.Atoi(rawProduct_id)
	// if err != nil {
	// 	return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	// }

	// err = h.ProductService.DeleteProduct(product_id)
	// if err != nil {
	// 	return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	// }

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "success to delete product", nil)
	httpresponse.Text(w, http.StatusOK, "Success")
}

func (h HttpHandlerImpl) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	// productRequestBody := new(dto.ProductsRequestBody)
	// if err := ctx.BodyParser(productRequestBody); err != nil {
	// 	return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	// }

	// product, err := h.ProductService.CreateNewProduct(*productRequestBody)
	// if err != nil {
	// 	// logger.Log.Errorln(err)
	// 	return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	// }

	// productResponse := dto.ParseFromEntity(product)

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", productResponse)
	httpresponse.Text(w, http.StatusOK, "Success")
}
