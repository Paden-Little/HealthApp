package security

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const jwtSecret = "supertopsecrethealthmark"

type Claims struct {
	ProviderID string `json:"provider_id"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for a given provider ID
func GenerateJWT(providerID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ProviderID: providerID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return tokenString, nil
}

// AuthMiddleware is a middleware that checks the JWT token in the Authorization header
func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		c.Abort()
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	if !token.Valid || claims.ProviderID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	if claims.ProviderID != c.Param("id") {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
		c.Abort()
		return
	}

	c.Next()
}
