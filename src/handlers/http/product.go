package http

import (
	httpresponse "golang-starter/internal/protocols/http/response"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h HttpHandlerImpl) GetProducts(ctx *fiber.Ctx) error {
	// get all products
	// products := h.ProductService.GetProducts()

	// productsResponse := dto.ParseFromEntities(products)

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", productsResponse)
	return httpresponse.TextResponse(ctx, http.StatusOK, "OK")
}

func (h HttpHandlerImpl) GetProductByID(ctx *fiber.Ctx) error {
	rawProduct_id := ctx.Params("product_id")
	product_id, err := strconv.Atoi(rawProduct_id)
	if err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	// product := h.ProductService.GetProductByProductID(product_id)

	// productResponse := dto.ParseFromEntity(product)

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", productResponse)
	return httpresponse.TextResponse(ctx, http.StatusOK, product_id)
}

func (h HttpHandlerImpl) DeleteProductByID(ctx *fiber.Ctx) error {
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
	return httpresponse.TextResponse(ctx, http.StatusOK, "Success")
}

func (h HttpHandlerImpl) CreateNewProduct(ctx *fiber.Ctx) error {
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
	return httpresponse.TextResponse(ctx, http.StatusOK, "Success")
}
