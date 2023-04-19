package main

import (
	"fmt"
	"grfc/controller"
	"grfc/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectToDatabase();
	database.SetUpRedis();
	app := fiber.New();
	app.Get("/user",controller.GetUsers);
	app.Post("/user",controller.CreateUser);
	pong, err := database.Cl.Ping().Result()
	fmt.Println(pong, err)
	log.Fatal(app.Listen(":3000"));
}