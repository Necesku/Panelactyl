package auth

import "github.com/joho/godotenv";
import "github.com/golang-jwt/jwt";
import "golang.org/x/crypto/bcrypt";
import "fmt";
import "os";
import "teletvbis/panelactyl/backend/database";

func GetFromEnv(key string) string {
	err := godotenv.Load(".env");
	if err != nil {
		fmt.Println("An error happened while loading env file!");
	}
	return os.Getenv(key);
}

func CreateToken(inputUser string) (string, error) {
	err, user := database.FindUser(inputUser);
	if err != nil {
		return "", err;
	}
	jwt_secret := GetFromEnv("JWT_SECRET");
	token_raw := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Name,
		"password": user.Password,
		"authorized": true,
	});
	token, err := token_raw.SignedString([]byte(jwt_secret));
	if err != nil {
		return "", err;
	}
	return token, nil;
}

func Login(user string, inputPassword string) (error, string) {
	err := Compare(inputPassword, user);
	if err != nil {
		return err, "";
	}
	token, err := CreateToken(user);
	if err != nil {
		return err, ""
	}
	return nil, token;
}

func Hash(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8); return string(hashed);
}

func Compare(inputPassword string, inputUser string) error {
	err, user := database.FindUser(inputUser);
	if err != nil {
		return err;
	}
	password := user.Password;
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(inputPassword));
	if err != nil {
		return err;
	}
	return nil;
}

func GetUserFromToken(tokenString string) (jwt.MapClaims, error) {
	jwt_secret := GetFromEnv("JWT_SECRET");
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