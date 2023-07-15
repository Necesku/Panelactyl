package main

import "fmt";
import "github.com/gofiber/fiber/v2";
import "github.com/joho/godotenv";
import "github.com/golang-jwt/jwt"
import "os";
import "golang.org/x/crypto/bcrypt";

func getFromEnv(key string) string {
	err := godotenv.Load("config.env");
	if err != nil {
		fmt.Println("An error happened while loading env file!");
	}
	return os.Getenv(key);
}

func createToken(user string, password string) (string, error) {
	jwt_secret := getFromEnv("ENV_SECRET");
	token_raw := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"authorized": true,
	});
	claims := token_raw.Claims.(jwt.MapClaims);
	claims["user"] = user;
	claims["password"] = password;
	claims["authorized"] = true;
	token, err := token_raw.SignedString([]byte(jwt_secret));
	if err != nil {
		return "", err;
	}
	return token, nil;
}

func hash(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8); return string(hashed);
}

func getUserFromToken(tokenString string) (jwt.MapClaims, error) {
	jwt_secret := getFromEnv("ENV_SECRET");
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method!");
		}
		return []byte(jwt_secret), nil;
	});
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil;
	} else {
		return nil, err;
	}
}

func main() {

	// JWT (createToken, getUserFromToken), HASHING (hash) TEST

	token, err := createToken("test", hash("hello!"));
	if err != nil {
		panic(err);
	}

	claims, err := getUserFromToken(token);
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