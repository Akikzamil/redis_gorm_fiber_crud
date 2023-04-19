package controller

import (
	"encoding/json"
	"grfc/database"
	"grfc/model"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Create(&user)
	getUsersFromDB()
	return c.Status(201).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := getUsersFromRedis()
	if err != nil {
		log.Println("jfeifiefj")
		users = getUsersFromDB()
	}
	return c.JSON(users)
}

func getUsersFromDB() []model.User {
	var users []model.User

	database.DB.Find(&users)
	cacheUsers(users)
	return users
}

func getUsersFromRedis() ([]model.User, error) {
	var users []model.User
	val, err := database.Cl.Get("users").Bytes()
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(val, &users)
	return users, err
}

func cacheUsers(users []model.User) {
	json, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	err2 := database.Cl.Set("users", json, time.Hour)
	if err != nil {
		log.Println(err2)
	}
}
