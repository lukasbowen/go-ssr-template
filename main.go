package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/recover"
  "github.com/gofiber/template/mustache/v2"
  "github.com/lukasbowen/go-ssr-template/handlers"
  "log"
)

func main() {
  fmt.Println("Hello World")

  engine := mustache.New("./views", ".mustache")
  engine.Reload(true)

  app := fiber.New(fiber.Config{
    Views: engine,
    ErrorHandler: func(ctx *fiber.Ctx, err error) error {
      return ctx.Render("errors/500", fiber.Map{
        "Title": "Server Error",
      }, "layouts/main")
    },
  })
  app.Use(recover.New())
  app.Static("/static", "./static")

  handlers.AuthRoutes(app)

  app.Get("/", func(c *fiber.Ctx) error {
    // Render index
    return c.RedirectToRoute("Login", fiber.Map{})
  })

  app.Get("/layout", func(c *fiber.Ctx) error {
    // Render index within layouts/main
    return c.Render("index", fiber.Map{
      "Title": "Hello, World!",
    }, "layouts/main")
  })

  log.Fatal(app.Listen(":5431"))
}
