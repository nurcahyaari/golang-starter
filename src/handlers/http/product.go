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

func (h HttpHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get all products
	products, err := h.ProductService.GetProducts(r.Context())
	if err != nil {
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", products)
}

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
	// productRequestBody := new(dto.ProductsRequestBody)
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
