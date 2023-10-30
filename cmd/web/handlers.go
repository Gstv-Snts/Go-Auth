package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)
type onChange struct {
    Change string `json:"change" xml:"change" form:"change"`
}

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", nil)
}

func HomePage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func OnChange(c *fiber.Ctx) error {
    changeBody := new(onChange)
    err := c.BodyParser(changeBody) 
    if err != nil{
        fmt.Printf("Error parsing body: %v", err)
    }
    runes := []rune(changeBody.Change)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    c.Status(200)
	return c.SendString(string(runes))
}
