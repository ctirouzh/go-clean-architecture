package auth

import (
	"fmt"
	"lms/config"
	"lms/internal/domain/user"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	ID       string `json:"id"`
	Username string `json:"username"`
	Type     string `json:"type"`
}

// JwtManager is a json web token manager.
type JwtManager struct {
	sercetKey string
	ttl       time.Duration
}

// NewJwtManager returns a new JWT mananer using given JWT config.
func NewJwtManager(config config.JWT) *JwtManager {
	return &JwtManager{
		sercetKey: config.SecretKey,
		ttl:       config.TTL,
	}
}

// Generate generates and signs new access token for a user.
func (m *JwtManager) Generate(usr *user.User) (string, error) {

	if !usr.IsVerified() {
		return "", user.ErrUserNotVerified
	}
	if usr.IsBanned() {
		return "", user.ErrBannedUser
	}

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID:       usr.ID.String(),
		Username: usr.Username,
		Type:     usr.Type.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.sercetKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (m *JwtManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return []byte(m.sercetKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
