package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {

	app := fiber.New(fiber.Config{
		//Immutable: true, //Tüm requestleri verileri kalıcı durumda olur
		//Prefork: true, // Bir tane ana process oluşur çekirdek sayısı kadar childs üretir go ilk setup aşamasında
	})

	/*if fiber.IsChild() == false { //ise ana process demektir child yönetir else ise child process dir child process de ilgili akışı yürütür. Aynı portu dinler
		fmt.Println("Parent process started")
	}*/

	app.Use(func(c *fiber.Ctx) error { //tüm requestleri yakalama
		fmt.Println("Middleware 1")
		return c.Next() //Diğer handlerları kontrol etmesi için
	}, func(c *fiber.Ctx) error { //handlers birden fazla tanımlanarak devam edilebiir
		fmt.Println("Middleware 2")
		return c.Next()
	})

	app.Static("/", "./public") // Dizin paylaşımı
	app.Static("/assets", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test-error", func(c *fiber.Ctx) error {
		return fiber.NewError(400, "Custom error message")
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

	api := app.Group("/v1") // Grup tanımlama

	api.Get("/accounts", func(c *fiber.Ctx) error { // Gruba route tanımlama
		return c.SendString("accounts")
	})

	app.Route("/v3", func(api fiber.Router) { //Gruplama
		api.Get("/product", func(c *fiber.Ctx) error { //Gruba ait route tanımlama
			return c.SendString("product")
		}).Name("product") //Route Name Tanımlama
	}, "products") //Grup name tanımlama

	app.Listen(":3000")
}
