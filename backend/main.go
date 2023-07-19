package main

import "fmt";
import "errors";
import "github.com/gofiber/fiber/v2";
import "teletvbis/panelactyl/backend/auth";
import "teletvbis/panelactyl/backend/database";
import "database/sql";
import "golang.org/x/crypto/bcrypt";

type LoginBody struct {
	Username string;
	Password string;	
}

func main() {

	port := 8080;
	app := fiber.New();

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🟢 Pong!");
	});

	app.Post("/user/auth/login", func(c *fiber.Ctx) error {
		body := new(LoginBody);
		err := c.BodyParser(body);
		if err != nil {
			return c.JSON(fiber.Map{"error": true, "message": "BodyParser error"});
		}
		err, token := auth.Login(body.Username, body.Password);
		switch {
			case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
				return c.JSON(fiber.Map{"error": true, "message": "Error: Given password is wrong!"});
			case errors.Is(err, sql.ErrNoRows):
				return c.JSON(fiber.Map{"error": true, "message": "Error: Given username is wrong!"});
			default:
				return c.JSON(fiber.Map{"error": true, "message": "Error: Unknown error."});
			case err == nil:
				return c.JSON(fiber.Map{"error": false, "token": token});
		}
	});

	app.Post("user/auth/register", func(c* fiber.Ctx) error {
		body := new(LoginBody);
		err := c.BodyParser(body);
		if err != nil {
			return c.JSON(fiber.Map{"error": true, "message": "BodyParser error"});
		}
		err = auth.Register(body.Username, body.Password);
		switch {
			case errors.Is(err, database.ErrUserExists):
				return c.JSON(fiber.Map{"error": true, "message": "Error: Account with that name already exists!"});
			default:
				return c.JSON(fiber.Map{"error": true, "message": "Error: Unknown error."});
			case err == nil:
				return c.JSON(fiber.Map{"error": false, "message": "Authorize again!"});
		}
	});

	app.Listen(fmt.Sprintf(":%d", port));

}