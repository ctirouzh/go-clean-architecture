package auth

import (
	"lms/config"
	"lms/internal/domain/user"
	"lms/internal/pkg/sample"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_NewService(t *testing.T) {
	// mock domain user.Repository Interface implementation
	userRepo := struct {
		user.Repository
	}{}
	jwtManager := NewJwtManager(config.JWT{SecretKey: "secret", TTL: time.Minute})

	authService := NewService(userRepo, jwtManager)
	assert.NotEmpty(t, authService)
}

func TestAuthService_SignUp(t *testing.T) {
	// use domain user.Repository Interface mock implementation
	userRepo := user.NewMockRepository(make(map[uuid.UUID]*user.User))
	jwtManager := NewJwtManager(config.JWT{SecretKey: "secret", TTL: time.Minute})
	authService := NewService(userRepo, jwtManager)

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

	_, findErr := authService.userRepo.Get(usr.ID)
	assert.Empty(t, findErr)

}

func TestAuthService_SignIn(t *testing.T) {
	// use domain user.Repository Interface mock implementation
	users := make(map[uuid.UUID]*user.User)
	usr := sample.NewFakeUserEntity(user.USER_TYPE_ADMIN, true, false)
	users[usr.ID] = &usr
	userRepo := user.NewMockRepository(users)
	jwtManager := NewJwtManager(config.JWT{SecretKey: "secret_key", TTL: time.Minute})
	authService := NewService(userRepo, jwtManager)

	testCases := []struct {
		name     string
		username string
		password string
		want     error
	}{
		{
			name:     "valid username and password",
			username: usr.Username,
			password: "secret", // NewFakeUserEntity set password to "secret" internally.
			want:     nil,
		},
		{
			name:     "invalid username",
			username: "john_doe",
			password: "secret",
			want:     user.ErrUserNotFound,
		},
		{
			name:     "invalid password",
			username: usr.Username,
			password: "invalid",
			want:     ErrInvalidPassword,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, signInErr := authService.SignIn(tc.username, tc.password)
			assert.Equal(t, tc.want, signInErr)
			if signInErr == nil {
				assert.NotEmpty(t, token)
				_, verErr := authService.jwtManager.Verify(token)
				assert.Empty(t, verErr)
			}
		})
	}
}
