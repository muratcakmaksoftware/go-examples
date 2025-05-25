package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"` //interace == any
}

type ErrorResponse struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
}

func Success(c *fiber.Ctx, message string, data any) error {
	return JsonResponse(c, Response{
		Message: message,
		Data:    data,
	}, fiber.StatusOK, message, data)
}

func Error(c *fiber.Ctx, errorCode int, message string, data any) error {
	return JsonResponse(c, ErrorResponse{
		Error: message,
		Data:  data,
	}, errorCode, message, data)
}

func JsonResponse(c *fiber.Ctx, model any, status int, message string, data any) error {
	return c.Status(status).JSON(Response{
		Message: message,
		Data:    data,
	})
}

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

	app.Get("/stack", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.JSON(c.App().Stack())
	})

	app.Get("/sample-response", func(c *fiber.Ctx) error {
		user := User{
			Name: "Murat",
			Role: "admin",
		}
		return Success(c, "Kullanıcı getirildi", user)
	})

	app.Get("/sample-error-response", func(c *fiber.Ctx) error {
		return Error(c, fiber.StatusBadRequest, "Kötü İstek", nil)
	})

	app.Get("/profile/:murat", func(c *fiber.Ctx) error { //Parametre alimi immutable değil. "Handler içindeysek yapmamız gerekli değil."
		params := c.AllParams() // tüm parametreleri almak için
		return c.JSON(fiber.Map{
			"params":  params,
			"message": "Merhaba, " + c.Params("murat") + "!",
		})
	})

	app.Get("/deep-copy/:foo", func(c *fiber.Ctx) error {
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
