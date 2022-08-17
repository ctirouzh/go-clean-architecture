package controller

import (
	"bytes"
	"encoding/json"
	"lms/config"
	"lms/internal/adapters/repository/memory"
	"lms/internal/app/auth"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_SignUp(t *testing.T) {
	// Prepare auth controller dependencies.
	userRepo := memory.NewUserRepo()
	jwtManager := auth.NewJwtManager(config.JWT{SecretKey: "secret", TTL: time.Minute})
	authService := auth.NewService(userRepo, jwtManager)
	authCtrl := NewAuthController(authService)
	// testCases: TODO--> Consider EmailAlreadyTaken and UsernameAlreadyTaken errors
	testCases := []struct {
		name string
		form SignUpRequest
		want int
	}{
		{
			name: "valid signup form",
			form: SignUpRequest{
				Username: "test_username",
				Email:    "test@test.io",
				Password: "secret",
			},
			want: http.StatusOK,
		},
		{
			name: "empty username in signup form",
			form: SignUpRequest{
				Username: "",
				Email:    "test@test.io",
				Password: "secret",
			},
			want: http.StatusBadRequest,
		},
		{
			name: "empty email in signup form",
			form: SignUpRequest{
				Username: "test_username",
				Email:    "",
				Password: "secret",
			},
			want: http.StatusBadRequest,
		},
		{
			name: "empty password in signup form",
			form: SignUpRequest{
				Username: "test_username",
				Email:    "test@test.io",
				Password: "",
			},
			want: http.StatusBadRequest,
		},
		{
			name: "invalid email in signup form",
			form: SignUpRequest{
				Username: "test_username",
				Email:    "test",
				Password: "secret",
			},
			want: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare context
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			jsonForm, _ := json.Marshal(tc.form)
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonForm))

			authCtrl.SignUp(ctx)
			assert.Equal(t, tc.want, w.Code)

			if w.Code == http.StatusOK {
				var got map[string]UserDTO
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					/* can also use: json.NewDecoder(w.Body).Decode(&got) */
					t.Fatal(err)
				}

				id, uuidErr := uuid.Parse(got["data"].ID)
				if uuidErr != nil {
					t.Fatal(uuidErr)
				}
				if _, err := userRepo.Get(id); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.form.Username, got["data"].Username)
				assert.Equal(t, tc.form.Email, got["data"].Email)
			}
		})
	}
}
