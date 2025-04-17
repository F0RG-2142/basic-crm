package main

import (
	"fmt"

	"github.com/F0RG-2142/basic-crm/database"
	"github.com/F0RG-2142/basic-crm/lead"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/vi/lead/:id", lead.DeleteLead)
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDB()
	db, err := database.DBConn.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	setupRoutes(app)
	app.Listen(8080)
}
