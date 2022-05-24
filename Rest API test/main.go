package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type shoppingList struct {
	UID       string  `json:"uid"`
	Item      string  `json:"item"`
	Total     float64 `json:"total"`
	Successed bool    `json:"successed"`
}

var orders = []shoppingList{
	{UID: "1", Item: "Banana", Total: 500, Successed: false},
	{UID: "2", Item: "Egg", Total: 1000, Successed: true},
	{UID: "3", Item: "Pomelo", Total: 700, Successed: false},
}

func getOders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, orders)
}

func postOrders(context *gin.Context) {
	var newOrder shoppingList

	if err := context.BindJSON(&newOrder); err != nil {
		return
	}

	orders = append(orders, newOrder)

	context.IndentedJSON(http.StatusCreated, newOrder)
}

func main() {
	router := gin.Default()
	router.GET("/orders", getOders)
	router.POST("/orders", postOrders)
	router.Run("localhost:9090")
}
