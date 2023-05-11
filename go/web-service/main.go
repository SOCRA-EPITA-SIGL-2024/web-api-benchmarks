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
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type userInput struct {
	Position position `json:"position"`
}

type product struct {
	CategoryId string  `json:"categoryId"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}

type garden struct {
	Title    string    `json:"title"`
	Id       string    `json:"id"`
	Position position  `json:"position"`
	Products []product `json:"products"`
}

func loadDataset() []garden {
	gardenJsonFilePath := os.Getenv("GARDENS_JSON_FILE_PATH")
	content, err := os.ReadFile(gardenJsonFilePath)
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
	userInput := userInput{}
	radiusParam, exist := ctx.GetQuery("radius")
	if !exist {
		fmt.Println("params missing")
		ctx.Status(http.StatusBadRequest)
		return
	}

	radius, err := strconv.Atoi(radiusParam)
	if err != nil {
		fmt.Println("convertion params failed")
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = ctx.ShouldBindJSON(&userInput)
	if err != nil {
		fmt.Println("failed bind json:", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	var products []product

	for _, garden := range gardens {
		products = append(products, garden.checkDistance(userInput.Position, radius)...)
	}
	ctx.JSON(http.StatusOK, products)
}

func main() {

	// load json
	fmt.Println("Loading json")
	loadDataset()
	fmt.Println("finish Loading json, key: ", len(gardens))

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
