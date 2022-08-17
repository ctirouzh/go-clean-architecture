package auth

import (
	"errors"
	"lms/config"
	"lms/internal/domain/user"
	"lms/internal/pkg/sample"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwtManager_New(t *testing.T) {

	cfg := config.JWT{SecretKey: "secret", TTL: 20 * time.Minute}

	got := NewJwtManager(cfg)
	assert.NotEmpty(t, got)
	assert.Equal(t, cfg.SecretKey, got.sercetKey)
	assert.Equal(t, cfg.TTL, got.ttl)
}

func TestJWTManager_Generate(t *testing.T) {
	cfg := config.JWT{SecretKey: "secret", TTL: 20 * time.Minute}

	jwtManager := NewJwtManager(cfg)

	testCases := []struct {
		name          string
		user          user.User
		wantGenErr    error
		wantVerHasErr bool
	}{
		{
			name:          "verified and permitted user",
			user:          sample.NewFakeUserEntity(user.USER_TYPE_ADMIN, true, false),
			wantGenErr:    nil,
			wantVerHasErr: false,
		}, {
			name:          "unverified user",
			user:          sample.NewFakeUserEntity(user.USER_TYPE_ADMIN, false, false),
			wantGenErr:    user.ErrUserNotVerified,
			wantVerHasErr: true,
		}, {
			name:          "verified/banned user",
			user:          sample.NewFakeUserEntity(user.USER_TYPE_ADMIN, true, true),
			wantGenErr:    user.ErrBannedUser,
			wantVerHasErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accessToken, gotGenErr := jwtManager.Generate(&tc.user)
			assert.Equal(t, tc.wantGenErr, gotGenErr)
			_, gotVerErr := jwtManager.Verify(accessToken)
			assert.Equal(t, tc.wantVerHasErr, gotVerErr != nil)
		})
	}
}

func TestJWTManager_Verify(t *testing.T) {

	cfg := config.JWT{SecretKey: "secret", TTL: 20 * time.Minute}

	jwtManager := NewJwtManager(cfg)
	usr := sample.NewFakeUserEntity(user.USER_TYPE_TEACHER, true, false)
	token, err := jwtManager.Generate(&usr)
	if err != nil {
		t.Fatal(err)
	}
	usrClaims, verErr := jwtManager.Verify(token)
	assert.Empty(t, verErr)
	if verErr == nil {
		assert.Equal(t, usr.ID.String(), usrClaims.ID)
		assert.Equal(t, usr.Username, usrClaims.Username)
	}

	// generate a malicious token!
	malConfig := config.JWT{SecretKey: "malicious", TTL: 2000 * time.Minute}
	malJwtTool := NewJwtManager(malConfig)
	malUser := usr
	malUser.Type = user.USER_TYPE_ADMIN
	malToken, malErr := malJwtTool.Generate(&malUser)
	if malErr != nil {
		t.Fatal(malErr)
	}
	_, malVerErr := jwtManager.Verify(malToken)
	assert.NotEmpty(t, malVerErr)
	assert.Equal(t, errors.New(ErrInvalidToken.Error()+": "+"signature is invalid"), malVerErr)
}
