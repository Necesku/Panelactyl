package auth

import "github.com/joho/godotenv";
import "github.com/golang-jwt/jwt";
import "golang.org/x/crypto/bcrypt";
import "fmt";
import "os";

func GetFromEnv(key string) string {
	err := godotenv.Load("config.env");
	if err != nil {
		fmt.Println("An error happened while loading env file!");
	}
	return os.Getenv(key);
}

func CreateToken(user string, password string) (string, error) {
	jwt_secret := GetFromEnv("ENV_SECRET");
	token_raw := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"password": password,
		"authorized": true,
	});
	token, err := token_raw.SignedString([]byte(jwt_secret));
	if err != nil {
		return "", err;
	}
	return token, nil;
}

func Hash(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8); return string(hashed);
}

func GetUserFromToken(tokenString string) (jwt.MapClaims, error) {
	jwt_secret := GetFromEnv("ENV_SECRET");
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