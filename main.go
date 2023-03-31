package main

import (
	"net/http"
	"strconv"
	"github.com/angelaguirremeli/get-method-live-class/products"
	"github.com/gin-gonic/gin"
)

var productsSlice, errFile = products.LoadFromJsonFile("products.json")

func main()  {
	
	if errFile != nil {
		panic(errFile)
	}

	router := Router()

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, productsSlice)
	})

	router.GET("/products/:id",  func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		} 
		
		idNumber, err := strconv.Atoi(id)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		}

		product := products.SearchProductById(idNumber, productsSlice)

		if product == nil {
			ctx.String(http.StatusNotFound, "Not Found")
			return
		} 
		ctx.JSON(http.StatusOK, product)
	})

	router.GET("/products/search", func(ctx *gin.Context) {
		price := ctx.Query("price")

		if price == "" {
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		} 

		priceAsNumber, err := strconv.ParseFloat(price, 64)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")
			return
		}

		productsFiltered := products.SearchProductByPrice(priceAsNumber, productsSlice)
		ctx.JSON(http.StatusOK, productsFiltered)
	})

	return router
}