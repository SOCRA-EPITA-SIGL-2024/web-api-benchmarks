package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

var gardens []garden

type position struct {
	lat float64 `json:"lat"`
	lng float64 `json:"lng"`
}

type product struct {
	categoryId string  `json:"categoryId"`
	name       string  `json:"name"`
	price      float64 `json:"price"`
}

type garden struct {
	title    string    `json:"title"`
	id       string    `json:"id"`
	position position  `json:"position"`
	products []product `json:"products"`
}

func loadDataset() []garden {

	content, err := os.ReadFile("./gardens.json")
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	json.Unmarshal(content, &gardens)

	return gardens
}

func ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "{\"Hello\": \"Socarotte\"}")
}

func noRoute(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

/**
 * Find all products that are 10km for the position of the user
 */
func allProductAround(ctx *gin.Context) {
	var userPosition position
	radiusParam, exist := ctx.GetQuery("radius")
	if !exist {
		ctx.Status(http.StatusBadRequest)
		return
	}

	radius, err := strconv.Atoi(radiusParam)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = ctx.ShouldBindJSON(userPosition)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var products []product

	for _, garden := range gardens {
		products = append(products, garden.checkDistance(userPosition, radius)...)
	}
	ctx.JSON(http.StatusOK, products)
}

func main() {

	// load json
	loadDataset()

	r := gin.New()
	r.NoRoute(noRoute)
	r.GET("/ping", ping)
	v1 := r.Group("/v1")
	{
		v1.POST("/products", allProductAround)
	}

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
