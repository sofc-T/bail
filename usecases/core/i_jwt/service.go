// Package ijwt provides JWT generation and decoding services.
package ijwt

import (
	"bail/domain/models"

	"github.com/dgrijalva/jwt-go"
)

const (
	Access  = "acccess"
	Refresh = "refresh"
	Reset   = "reset"
)

// Service defines methods to generate and decode JWTs.
type Service interface {
	// Generate creates a JWT for a user.
	Generate(user *models.User, tokenType string) (string, error)

	//Decode parses a JWT and returns claims in the claims it identifies whether user or businessOwner owner
	Decode(token string) (jwt.MapClaims, error)
}
