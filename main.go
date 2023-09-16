package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/redmicelles/Hotel-Reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "API Server Listening..")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleListUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
