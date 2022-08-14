package memory

import (
	"lms/internal/domain/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T) {
	username := "test_username"
	email := "test@test.io"
	password := "secret"
	userRepo := NewUserRepo()
	usr, createErr := userRepo.Create(username, email, password, user.USER_TYPE_STUDENT)
	if !assert.Equal(t, nil, createErr) {
		// username or email already taken
		t.Fatal("cannot create test user")
	}

	testCases := []struct {
		name     string
		id       uuid.UUID
		expected error
	}{
		{
			name:     "get and delete an existing user in repository",
			id:       usr.ID,
			expected: nil,
		},
		{
			name:     "try to find or delete a non-existent user",
			id:       uuid.New(),
			expected: user.ErrUserNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			found, getErr := userRepo.Get(tc.id)
			assert.Equal(t, tc.expected, getErr)
			if found != nil {
				assert.Truef(t, found.IsStudent(), "expected a student, got %s", found.Type.String())
				assert.False(t, found.IsVerified(), "expected an unverefied user, got a verified one")
				assert.False(t, found.IsBanned(), "expected a permitted user, got a banned one")
				assert.Equal(t, usr.Username, found.Username)
				assert.Equal(t, usr.Email, found.Email)
			}
			deleteErr := userRepo.Delete(usr.ID)
			assert.Equal(t, tc.expected, deleteErr)
		})
	}
}
