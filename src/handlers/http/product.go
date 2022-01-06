package http

import (
	"encoding/json"
	httpresponse "golang-starter/internal/protocols/http/response"
	"golang-starter/src/modules/product/dto"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

// GetProducts return Products list
// @Summary Get all products
// @Description get all products
// @Tags Products
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products [GET]
func (h HttpHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get all products
	products, err := h.ProductService.GetProducts(r.Context())
	if err != nil {
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", products)
}

// GetProductByID return Products by productId
// @Summary Get Products by productId
// @Description get Products by productId
// @Tags Products
// @Param productId path string true "productId"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/{productId} [GET]
func (h HttpHandlerImpl) GetProductByID(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "productId")
	productId, err := strconv.Atoi(rawProductId)
	if err != nil {
		httpresponse.Err(w, err)
		return
	}

	product, err := h.ProductService.GetProductByProductID(r.Context(), productId)
	if err != nil {
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", product)
}

func (h HttpHandlerImpl) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "productId")
	productId, err := strconv.Atoi(rawProductId)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	err = h.ProductService.DeleteProduct(r.Context(), productId)
	if err != nil {
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "success to delete product", nil)
}

func (h HttpHandlerImpl) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	product := dto.ProductRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	productNew, err := h.ProductService.CreateNewProduct(r.Context(), product)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "success create product", productNew)
}
