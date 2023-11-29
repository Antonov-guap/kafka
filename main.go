package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/samber/lo"
)

var panicOnErr = lo.Must0

func main() {
	app := fiber.New()
	app.Static("/", "views")
	app.Get("/chat", websocket.New(getChat))
	panicOnErr(app.Listen(":8081"))
}

func getChat(c *websocket.Conn) {
	for range time.Tick(1 * time.Second) {
		_ = c.WriteJSON(fiber.Map{
			"message": "hello",
			"number":  rand.Intn(1000),
		})
	}
}
