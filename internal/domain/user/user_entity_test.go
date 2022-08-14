package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity_CheckUserType(t *testing.T) {

	testCases := []struct {
		name     string
		user     User
		expected []bool
	}{
		{
			name:     "call {IsAdmin,IsTecher,IsStudent} on an admin user type",
			user:     User{Type: USER_TYPE_ADMIN},
			expected: []bool{true, false, false},
		}, {
			name:     "call {IsAdmin,IsTecher,IsStudent} on a teacher user type",
			user:     User{Type: USER_TYPE_TEACHER},
			expected: []bool{false, true, false},
		}, {
			name:     "call {IsAdmin,IsTecher,IsStudent} on a student user type",
			user:     User{Type: USER_TYPE_STUDENT},
			expected: []bool{false, false, true},
		}, {
			name:     "call {IsAdmin,IsTecher,IsStudent} on an unknown user type",
			user:     User{Type: -1},
			expected: []bool{false, false, false},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := []bool{tc.user.IsAdmin(), tc.user.IsTeacher(), tc.user.IsStudent()}
			assert.Equal(t, tc.expected, got)
		})
	}

}

func TestUserEntity_PrepareForCreate(t *testing.T) {
	usr := User{}
	usr.PrepareForCreate()
	_, uuidErr := uuid.Parse(usr.ID.String())
	assert.ErrorIs(t, nil, uuidErr)
	assert.False(t, usr.IsVerified(), "expected an unverified user, got a verified one")
	assert.False(t, usr.IsBanned(), "expected a permitted user, got a banned one")
	assert.False(t, usr.CreatedAt.IsZero())
	assert.False(t, usr.UpdatedAt.IsZero())
}

func TestUserEntity_SetPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		expected error
	}{
		{
			name:     "non-empty password",
			password: "secret",
			expected: nil,
		}, {
			name:     "empty password",
			password: "",
			expected: ErrEmptyUserPassword,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			usr := User{}
			gotErr := usr.SetPassword(tc.password)
			assert.Equal(t, tc.expected, gotErr)
			switch gotErr {
			case nil:
				assert.True(t, usr.IsPasswordVerified(tc.password))
				assert.False(t, usr.IsPasswordVerified("jskla9"))
			default:
				assert.False(t, usr.IsPasswordVerified(tc.password))
			}
		})
	}
}
