package main

import "fmt";
import "github.com/gofiber/fiber/v2";
import "teletvbis/panelactyl/backend/auth";

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
		if err != nil {
			if (err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password") {
				return c.JSON(fiber.Map{"error": true, "message": "Error: Given password is wrong!"});
			} else if (err.Error() == "sql: no rows in result set") {
				return c.JSON(fiber.Map{"error": true, "message": "Error: Given username is wrong!"});
			} else {
				return c.JSON(fiber.Map{"error": true, "message": "Error: Unknown error."});
			}
		}
		return c.JSON(fiber.Map{"error": false, "token": token});
	});

	app.Listen(fmt.Sprintf(":%d", port));

}