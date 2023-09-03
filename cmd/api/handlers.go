package main

import (
	"fmt"
	"net/http"
	"product-api/cmd/data"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func (app *Config) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	product := data.Product{
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
		Price:       requestPayload.Price,
		Status:      1,
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	}

	err = product.Create()

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Product created successfully by id %d", product.ID),
		Data:    product,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}

func (app *Config) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.Models.Product.GetAll()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "All products",
		Data:    products,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) GetProduct(w http.ResponseWriter, r *http.Request) {
	// get id from url as a path parameter
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	// convert id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	product, err := app.Models.Product.GetByID(intID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Product with id %d", intID),
		Data:    product,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get id from url as a path parameter
	id := chi.URLParam(r, "id")
	// convert id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var requestPayload struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	product := data.Product{
		ID:          intID,
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
		Price:       requestPayload.Price,
		Status:      1,
		UpdatedOn:   time.Now(),
	}

	err = product.Update()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Product updated successfully by id %d", product.ID),
		Data:    product,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id from url as a path parameter
	id := chi.URLParam(r, "id")
	// convert id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	product := data.Product{
		ID: intID,
	}

	err = product.Delete()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Product deleted successfully by id %d", product.ID),
		Data:    product,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) UndeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get id from url as a path parameter
	id := chi.URLParam(r, "id")
	// convert id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	product := data.Product{
		ID: intID,
	}

	err = product.Undelete()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Product undeleted successfully by id %d", product.ID),
		Data:    product,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
