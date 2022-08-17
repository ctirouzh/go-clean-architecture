package auth

import (
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
			name:          "verified/permitted user",
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

	testCases := []struct {
		name       string
		token      string
		wantHasErr bool
	}{
		{
			name:       "empty access token",
			token:      "",
			wantHasErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, gotErr := jwtManager.Verify(tc.token)
			assert.Equal(t, tc.wantHasErr, gotErr != nil)
			//TODO: check some details of got claims...
		})
	}
}
