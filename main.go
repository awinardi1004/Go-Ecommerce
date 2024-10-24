package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	// Rute CRUD Produk
	router.GET("/products", handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))
	
	// Rute CRUD Product Category
	router.GET("/product-categories", handlers.ListProductCategories(db))
	router.GET("/product-categories/:id", handlers.GetProductCategory(db))
	router.POST("/product-categories", handlers.CreateProductCategory(db))
	router.PUT("/product-categories/:id", handlers.UpdateProductCategory(db))
	router.DELETE("/product-categories/:id", handlers.DeleteProductCategory(db))

	router.POST("/transactions", handlers.CreateTransaction(db))
	router.GET("/transactions/:id", handlers.GetTransactionWithItems(db))

	



	router.Run(":8080")