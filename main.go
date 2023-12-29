package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	storage, err := NewMysqlStorage()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/triggerFileRead", func(c *fiber.Ctx) error {
		data, err := storage.ReadFromFile()
		if err != nil {
			return err
		}
		return c.JSON(data)
	})

	app.Listen(":3000")
}
