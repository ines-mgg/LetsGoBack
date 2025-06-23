package Context

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtsecret []byte

// verifyJWTSecret checks if the JWT secret is set.
// It returns an error if the secret is not set, which is necessary for JWT operations.
func verifyJWTSecret() error {
	if jwtsecret == nil {
		return errors.New("JWT secret is not set")
	}
	return nil
}

// trimJWT removes the "Bearer " prefix from a JWT token string if it exists.
// This is useful for standardizing the token format before validation or parsing.
// It checks if the token string starts with "Bearer " and removes it, returning the trimmed token.
// If the token string does not start with "Bearer ", it returns the original token string.
func trimJWT(tokenStr string) string {
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		return tokenStr[7:]
	}
	return tokenStr
}

// getJWTClaims extracts and validates the claims from a JWT token string.
// It verifies the JWT secret, trims the token string, and parses the token.
// If the token is valid, it returns the claims as jwt.MapClaims.
// If the token is invalid or the claims cannot be parsed, it returns an error.
// This function is used to retrieve the claims from a JWT token, which can include user information, roles, and other metadata.
// It ensures that the token is properly signed and has not expired.
// The claims are returned as a map, allowing easy access to individual claim values.
// It is important to call SetJWTSecret before using this function to ensure the JWT secret is set.
func getJWTClaims(tokenStr string) (jwt.MapClaims, error) {
	verifyJWTSecret()
	tokenStr = trimJWT(tokenStr)
	if tokenStr == "" {
		return nil, errors.New("token string is empty")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtsecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

// SetJWTSecret sets the JWT secret used for signing and validating JWT tokens.
// It should be called before generating or validating JWT tokens.
// The secret is a byte slice that is used to sign the JWT tokens.
// It is important to set the JWT secret before using any JWT-related functions,
// as it is required for both generating and validating tokens.
// The secret should be kept secure and not exposed publicly.
// It is recommended to use a strong, random secret for production applications to ensure the security of the JWT tokens.
func SetJWTSecret(secret string) {
	jwtsecret = []byte(secret)
}

// GenerateJWT creates a new JWT token with the provided claims.
// It uses the HS256 signing method and the JWT secret set by SetJWTSecret.
// The claims parameter is a map of key-value pairs that represent the claims to be included in the token.
// The function returns the signed token as a string.
// If the JWT secret is not set, it returns an error.
// This function is used to create JWT tokens for authentication and authorization purposes.
// It allows you to include custom claims such as user ID, roles, and other metadata in the token.
func GenerateJWT(claims map[string]any) (string, error) {

	if jwtsecret == nil {
		return "", errors.New("JWT secret is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString(jwtsecret)
}

// ValidateJWT validates a JWT token string and returns the claims if valid.
// It checks if the JWT secret is set, trims the token string, and parses the token.
// If the token is valid, it returns the claims as jwt.MapClaims.
// If the token is invalid or has expired, it returns an error.
// This function is used to verify the authenticity of a JWT token and extract its claims.
// It ensures that the token is properly signed and has not expired.
// The claims can include user information, roles, and other metadata.
// It is important to call SetJWTSecret before using this function to ensure the JWT secret is set.
func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	verifyJWTSecret()
	tokenStr = trimJWT(tokenStr)
	if tokenStr == "" {
		return nil, errors.New("token string is empty")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtsecret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, errors.New("invalid expiration time")
	}
	if exp != nil && exp.Time.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	iat, err := claims.GetIssuedAt()
	if err != nil {
		return nil, errors.New("invalid issued at time")
	}
	if iat != nil && iat.Time.After(time.Now()) {
		return nil, errors.New("token issued in the future")
	}
	return getJWTClaims(tokenStr)
}
