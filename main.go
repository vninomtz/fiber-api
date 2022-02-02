package main

import (
	"log"

	"github.com/NinoVictor/fiber-api/database"
	"github.com/NinoVictor/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to API REST")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
}
