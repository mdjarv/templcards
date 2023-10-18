package routes

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/mdjarv/templcards/components"
	"github.com/mdjarv/templcards/pkg/database"
	"github.com/mdjarv/templcards/pkg/models"
)

type Card struct {
	models.Card
	URL string
}

func Setup(app *fiber.App, db *database.Database) {
	app.Get("/", func(c *fiber.Ctx) error {
		cards, err := db.GetCards()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return HTML(c, components.Main(cards))
	})

	app.Post("/cards", func(c *fiber.Ctx) error {
		front := c.FormValue("front")
		back := c.FormValue("back")

		if front == "" || back == "" {
			return c.Status(400).SendString("front and back are required")
		}

		c.Context().Logger().Printf("front: %s", front)
		card, err := db.AddCard(models.Card{
			Front: strings.Split(front, "\n"),
			Back:  strings.Split(back, "\n"),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return HTML(c, components.Card(card))
	})

}

func HTML(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", "text/html")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}
