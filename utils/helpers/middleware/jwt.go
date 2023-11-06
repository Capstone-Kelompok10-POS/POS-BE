package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


func GenerateTokenAdmin(AdminID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"id": AdminID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
		"role": "Admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenCashier(CashierID uint) (string, error) {
    jwtSecret := []byte(os.Getenv("SECRET_KEY"))

    claims := jwt.MapClaims{
        "id":   CashierID,
        "exp":   time.Now().Add(time.Hour * 1).Unix(),
        "iat":   time.Now().Unix(),
		"role": "Cashier",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ExtractTokenAdminId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		AdminId := claims["id"].(float64)
		return AdminId
	}
	return 0
}


func ExtractTokenCashierId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
	CashierId := claims["id"].(float64)
		return CashierId
	}
	return 0
}