package auth

import (
	"lms/internal/domain/user"
	"lms/internal/pkg/sample"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_NewService(t *testing.T) {
	// mock domain user.Repository Interface implementation
	userRepo := struct {
		user.Repository
	}{}
	authService := NewService(userRepo)
	assert.NotEmpty(t, authService)
}

func TestAuthService_SignUp(t *testing.T) {
	// use domain user.Repository Interface mock implementation
	userRepo := user.NewMockRepository(make(map[uuid.UUID]*user.User))
	authService := NewService(userRepo)

	username := sample.NewFakeUsername()
	email := sample.NewFakeEmail()
	password := "secret"
	usr, err := authService.SignUp(username, email, password)

	assert.Empty(t, err)
	assert.NotEmpty(t, usr)
	assert.NotEmpty(t, usr.ID)
	assert.Equal(t, username, usr.Username)
	assert.Equal(t, email, usr.Email)
	assert.Equal(t, user.USER_TYPE_STUDENT, usr.Type)
	assert.True(t, usr.IsPasswordVerified(password))
}
