package handlers

import "github.com/gofiber/fiber/v2"

type loginRequest struct {
  Email    string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=6,max=32"`
}

type registrationData struct {
  Email    string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required,min=6,max=32"`
}

func AuthRoutes(app *fiber.App) {

  // GET
  app.Get("/login", func(c *fiber.Ctx) error {
    return c.Render("auth/login", fiber.Map{
      "Title": "Login",
    }, "layouts/main")
  }).Name("Login")

  app.Get("/signup", func(c *fiber.Ctx) error {
    return c.Render("auth/signup", fiber.Map{
      "Title": "Signup",
    }, "layouts/main")
  }).Name("SignUp")

  // POST
  app.Post("/api/login", func(ctx *fiber.Ctx) error {
    var loginRequest loginRequest
    err := ctx.BodyParser(loginRequest)

    if err != nil {
      return err
    }

    if loginRequest.Email != "lbowen" && loginRequest.Password != "password" {
      // Error
    }
    // success
    return nil
  })

}
