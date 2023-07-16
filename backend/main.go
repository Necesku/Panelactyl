package main

import "fmt";
import "github.com/gofiber/fiber/v2";
import "teletvbis/panelactyl/backend/auth";

func main() {

	// JWT (createToken, getUserFromToken), HASHING (hash) TEST

	token, err := auth.CreateToken("test", auth.Hash("hello!"));
	if err != nil {
		panic(err);
	}

	claims, err := auth.GetUserFromToken(token);
	if err != nil {
		panic(err);
	}

	fmt.Println(claims);

	port := 8080;
	app := fiber.New();

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🟢 Pong!");
	});

	app.Listen(fmt.Sprintf(":%d", port));

}