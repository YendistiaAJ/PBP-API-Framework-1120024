package controllers

import (
	"net/http"
	d "test_revel/app/db"
	m "test_revel/app/models"

	"github.com/revel/revel"
)

type Product struct {
	*revel.Controller
}

func (c Product) GetProductById() revel.Result {
	id := c.Params.Route.Get("id")
	var response m.ProductResponse
	db := d.Connect()
	defer db.Close()

	rows, err := db.Query(m.GetProductByIdSQL, id)
	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.ProductResponse{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	var product m.Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			c.Log.Errorf("Rows Scan Error: %v", err)
			response = m.ProductResponse{
				Status:      http.StatusBadRequest,
				ContentType: "Failed to retrieve product data",
			}
			return c.RenderJSON(response)
		}
	}

	response = m.ProductResponse{
		Status:      http.StatusOK,
		ContentType: "Success",
		Data:        product,
	}
	return c.RenderJSON(response)
}

func (c Product) GetProducts() revel.Result {
	var response m.ProductsResponse
	db := d.Connect()
	defer db.Close()

	rows, err := db.Query(m.GetProductSQL)
	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.ProductsResponse{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	var products []m.Product
	for rows.Next() {
		var product m.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			c.Log.Errorf("Rows Scan Error: %v", err)
			response = m.ProductsResponse{
				Status:      http.StatusBadRequest,
				ContentType: "Failed to retrieve products data",
			}
			return c.RenderJSON(response)
		} else {
			products = append(products, product)
		}
	}

	response = m.ProductsResponse{
		Status:      http.StatusOK,
		ContentType: "Success",
		Data:        products,
	}
	return c.RenderJSON(response)
}
