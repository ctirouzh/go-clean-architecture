package controller

import (
	"lms/internal/domain/user"
	"lms/internal/pkg/sample"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthDTO_PrepareUserDTO(t *testing.T) {
	usr := sample.NewFakeUserEntity(user.USER_TYPE_STUDENT, false, false)
	usrDTO := UserDTO{}
	usrDTO.Prepare(&usr)
	assert.Equal(t, usrDTO.ID, usr.ID.String())
	assert.Equal(t, usrDTO.Username, usr.Username)
	assert.Equal(t, usrDTO.Email, usr.Email)
	assert.Equal(t, usrDTO.Type, usr.Type.String())
	assert.Equal(t, usrDTO.Verified, usr.Verified)
	assert.Equal(t, usrDTO.Banned, usr.Banned)
	assert.Equal(t, usrDTO.CreatedAt, usr.CreatedAt)
	assert.Equal(t, usrDTO.UpdatedAt, usr.UpdatedAt)
}
