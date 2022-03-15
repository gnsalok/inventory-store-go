package main

import (
	"log"
	"net/http"

	"github.com/gnsalok/inventory-store-go/database"
	"github.com/gnsalok/inventory-store-go/product"
	"github.com/gnsalok/inventory-store-go/receipt"
	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
