package main

import (
	"github.com/Gstv-Snts/Go-Auth/cmd/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("../ui/html", ".html")
	app := fiber.New(fiber.Config{
		Prefork: true,
		Views:   engine,
	})
	app.Get("/", web.HomePage)
	app.Get("/login", web.LoginPage)
	app.Get("/register", web.RegisterPage)
    app.Post("/change", web.OnChange)
	app.Listen(":8080")
}
