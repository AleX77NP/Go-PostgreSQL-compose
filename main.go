package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	const (
		HOST = "database"
		PORT = 5432
	)

	dbUser, dbPwd, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")

	dsn :=  fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	HOST, PORT, dbUser, dbPwd, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON("This is api route.")
	})

	log.Fatal(app.Listen(":8080"))
}