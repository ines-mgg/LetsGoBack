package core

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("change-this-to-something-secure")

func GenerateJWT(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ParseJWT(tokenStr string) (string, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return "", errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("invalid claims")
    }

    userID, ok := claims["user_id"].(string)
    if !ok {
        return "", errors.New("user_id not found")
    }

    return userID, nil
}
