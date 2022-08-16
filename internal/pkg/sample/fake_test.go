package sample

import (
	"lms/internal/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFake_NewUserEntity(t *testing.T) {
	// banned admin!
	usr := NewFakeUserEntity(user.USER_TYPE_ADMIN, false, true)
	assert.NotEmpty(t, usr.ID)
	assert.True(t, usr.IsAdmin())
	assert.False(t, usr.IsVerified())
	assert.True(t, usr.IsBanned())
}
func TestFake_NewUsername(t *testing.T) {
	username := NewFakeUsername()
	assert.NotEmpty(t, username)
}

func TestFake_NewEmail(t *testing.T) {
	email := NewFakeEmail()
	assert.NotEmpty(t, email)
}
