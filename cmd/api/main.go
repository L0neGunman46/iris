package main

import (
	"log"
	"os"

	"github.com/L0neGunman46/urlShortApp/internal/database"
	"github.com/L0neGunman46/urlShortApp/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment: ", err)
	}

	config := &database.Config{
		Conn: os.Getenv("CONN_STR"),
	}
	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatal("Error connecting with the database: ", err)
	}

	err = server.MigrateURLs(db)
	if err != nil {
		log.Fatal("Could not migrate database with schema ", err)
	}

	routes := server.NewRepository(db)

	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
