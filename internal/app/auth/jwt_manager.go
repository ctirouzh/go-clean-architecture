package auth

import (
	"errors"
	"lms/config"
	"lms/internal/domain/user"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrUserNotVerified         = errors.New("user not verified")
	ErrBannedUser              = errors.New("user is banned")
	ErrUnexpectedSigningMethod = errors.New("unexpected token signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrInvalidTokenClaims      = errors.New("invalid token claims")
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
		return "", ErrUserNotVerified
	}
	if usr.IsBanned() {
		return "", ErrBannedUser
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
				return nil, ErrUnexpectedSigningMethod
			}
			return []byte(m.sercetKey), nil
		},
	)

	if err != nil {
		return nil, errors.New(ErrInvalidToken.Error() + ": " + err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, ErrInvalidTokenClaims
	}

	return claims, nil
}
