package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viveksisodia1108/gogingorm/database"
	"github.com/viveksisodia1108/gogingorm/handlers"
)

func main() {
	database.ConnectDb()

	app := gin.Default()

	app.GET("/transactions/account/:acc_id", handlers.ListTrasactionsByAccountAndDate)

	app.GET("/transactions/:tx_id", handlers.ListTrasactionsById)

	app.Run(":3000")
}
