package lead

import (
	"github.com/F0RG-2142/basic-crm/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(503).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead Succesfully Deleted")
}
