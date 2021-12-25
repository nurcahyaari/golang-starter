package http

import (
	httpresponse "golang-starter/internal/protocols/http/response"
	"golang-starter/src/domains/product/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h HttpHandlerImpl) GetProducts(ctx *fiber.Ctx) error {
	// get all products
	products := h.ProductService.GetProducts()

	productsResponse := dto.ParseFromEntities(products)

	return httpresponse.JsonResponse(ctx, http.StatusOK, "", productsResponse)
}

func (h HttpHandlerImpl) GetProductByID(ctx *fiber.Ctx) error {
	rawProduct_id := ctx.Params("product_id")
	product_id, err := strconv.Atoi(rawProduct_id)
	if err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	product := h.ProductService.GetProductByProductID(product_id)

	productResponse := dto.ParseFromEntity(product)

	return httpresponse.JsonResponse(ctx, http.StatusOK, "", productResponse)
}

func (h HttpHandlerImpl) DeleteProductByID(ctx *fiber.Ctx) error {
	rawProduct_id := ctx.Params("product_id")
	product_id, err := strconv.Atoi(rawProduct_id)
	if err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	err = h.ProductService.DeleteProduct(product_id)
	if err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	return httpresponse.JsonResponse(ctx, http.StatusOK, "success to delete product", nil)
}

func (h HttpHandlerImpl) CreateNewProduct(ctx *fiber.Ctx) error {
	productRequestBody := new(dto.ProductsRequestBody)
	if err := ctx.BodyParser(productRequestBody); err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	product, err := h.ProductService.CreateNewProduct(*productRequestBody)
	if err != nil {
		// logger.Log.Errorln(err)
		return httpresponse.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	productResponse := dto.ParseFromEntity(product)

	return httpresponse.JsonResponse(ctx, http.StatusOK, "", productResponse)
}
