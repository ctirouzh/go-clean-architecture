package memory

import (
	"lms/internal/domain/user"
	"testing"

	"github.com/google/uuid"
)

func TestUserRepo(t *testing.T) {
	username := "test_username"
	email := "test@test.io"
	password := "secret"
	userRepo := NewUserRepo()
	usr, createErr := userRepo.Create(username, email, password, user.USER_TYPE_STUDENT)
	if createErr != nil { // username or email already taken
		t.Fatal("failed to initialize the test")
	}

	type testCase struct {
		name     string
		id       uuid.UUID
		expected error
	}

	testCases := []testCase{
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
			if getErr != tc.expected {
				t.Errorf("expected %v, got %v error", tc.expected, getErr)
			}
			if found != nil {
				if !found.IsStudent() {
					t.Errorf("expected a student, got %s", found.Type.String())
				}
				if found.IsVerified() {
					t.Error("expected an unverefied user, got a verified one")
				}
				if found.IsBanned() {
					t.Error("expected a permitted user, got a banned one")
				}
				if found.Username != usr.Username {
					t.Errorf("expected %s, got %s", usr.Username, found.Username)
				}
				if found.Email != usr.Email {
					t.Errorf("expected %s, got %s", usr.Email, found.Email)
				}
			}
			if deleteErr := userRepo.Delete(usr.ID); deleteErr != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, deleteErr)
			}
		})
	}
}
