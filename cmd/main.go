package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcherin/hackaton/database"
	"github.com/rcherin/hackaton/handlers"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/transactions/account/:acc_id", handlers.ListTrasactionsByAccountAndDate)

	app.Get("/transactions/:tx_id", handlers.ListTrasactionsById)

	app.Listen(":3000")
}
