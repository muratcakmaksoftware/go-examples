package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {

	app := fiber.New(fiber.Config{
		//Immutable: true, //Tüm requestleri verileri kalıcı durumda olur
	})

	app.Static("/", "./public") // Dizin paylaşımı
	app.Static("/assets", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:murat", func(c *fiber.Ctx) error { //Parametre alimi immutable değil. "Handler içindeysek yapmamız gerekli değil."
		return c.SendString("Isim: " + c.Params("value"))
	})

	app.Get("/:foo", func(c *fiber.Ctx) error {
		result := utils.CopyString(c.Params("foo")) //utils de hazır immutable edecek CopyString metodu. "Handler dışına veri taşınacaksa gereklidir."

		return c.SendString(result) //return nil hatasız demektir.
	})

	app.Get("/optional/test/:name?", func(c *fiber.Ctx) error { // optional parametre için örnektir `?` ile yapılır.
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
		}
		return c.SendString("Where is john?")
	})

	app.Get("/api/*", func(c *fiber.Ctx) error { //Joker `*` karakteri ile yazılan tüm path bilgisi alınabilir.
		return c.SendString("API path: " + c.Params("*"))
	})

	app.Listen(":3000")
}
