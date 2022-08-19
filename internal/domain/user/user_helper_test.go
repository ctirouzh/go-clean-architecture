package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHelper_NewUser(t *testing.T) {

	usr := NewUser(USER_TYPE_STUDENT)

	assertNotEmptyUser(t, usr)
	assert.True(t, usr.IsStudent())
	assert.False(t, usr.IsVerified())
	assert.False(t, usr.IsBanned())
}

func TestUserHelper_NewVerifiedUser(t *testing.T) {

	usr := NewVerifiedUser(USER_TYPE_TEACHER)

	assertNotEmptyUser(t, usr)
	assert.True(t, usr.IsTeacher())
	assert.True(t, usr.IsVerified())
	assert.False(t, usr.IsBanned())
}

func TestUserHelper_NewBannedUser(t *testing.T) {

	usr := NewBannedUser(USER_TYPE_ADMIN)

	assertNotEmptyUser(t, usr)
	assert.True(t, usr.IsAdmin())
	assert.True(t, usr.IsVerified())
	assert.True(t, usr.IsBanned())
}

func assertNotEmptyUser(t *testing.T, usr User) {
	assert.NotEmpty(t, usr.ID.String())
	assert.NotEmpty(t, usr.Username)
	assert.NotEmpty(t, usr.Email)
	assert.NotEmpty(t, usr.PasswordHash)
	assert.NotEmpty(t, usr.CreatedAt)
	assert.NotEmpty(t, usr.UpdatedAt)
}
