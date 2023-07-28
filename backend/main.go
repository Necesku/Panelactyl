package main

import "fmt";
import "errors";
import "github.com/gofiber/fiber/v2";
import "teletvbis/panelactyl/backend/auth";
import "teletvbis/panelactyl/backend/database";
import "database/sql";
import "golang.org/x/crypto/bcrypt";
import "github.com/gofiber/fiber/v2/middleware/cors";

type LoginBody struct {
	Username string `json:"username" xml:"username" form:"username"`;
	Password string `json:"password" xml:"password" form:"password"`;	
}

func main() {

	port := 3000;
	app := fiber.New();

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Access-Control-Allow-Origin, Content-Type",
		AllowOrigins: auth.GetFromEnv("FRONTEND_URL"),
	}));

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
				return c.JSON(fiber.Map{"error": true, "message": "Given password is wrong!"});
			case errors.Is(err, sql.ErrNoRows):
				return c.JSON(fiber.Map{"error": true, "message": "Given username is wrong!"});
			default:
				return c.JSON(fiber.Map{"error": true, "message": "Unknown error."});
			case err == nil:
				return c.JSON(fiber.Map{"error": false, "token": token});
		}
	});

	app.Post("/user/auth/register", func(c* fiber.Ctx) error {
		body := new(LoginBody);
		err := c.BodyParser(body);
		if err != nil {
			return c.JSON(fiber.Map{"error": true, "message": "Wrong input! (BodyParser error)"});
		}
		err = auth.Register(body.Username, body.Password);
		switch {
			case errors.Is(err, database.ErrUserExists):
				return c.JSON(fiber.Map{"error": true, "message": "Account with that name already exists!"});
			default:
				return c.JSON(fiber.Map{"error": true, "message": "Unknown error."});
			case err == nil:
				return c.JSON(fiber.Map{"error": false, "message": "Authorize again!"});
		}
	});

	app.Listen(fmt.Sprintf(":%d", port));

}