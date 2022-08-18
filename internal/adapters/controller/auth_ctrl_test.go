package controller

import (
	"bytes"
	"encoding/json"
	"lms/config"
	"lms/internal/adapters/repository/memory"
	"lms/internal/app/auth"
	"lms/internal/domain/user"
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
	jwtManager := auth.NewJwtManager(config.JWT{SecretKey: "secret_key", TTL: time.Minute})
	authService := auth.NewService(userRepo, jwtManager)
	authCtrl := NewAuthController(authService)
	// Test Cases: TODO--> Consider EmailAlreadyTaken and UsernameAlreadyTaken errors
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

func TestAuthController_SignIn(t *testing.T) {
	// use domain user.Repository Interface mock implementation
	validUser := user.NewVerifiedUser(user.USER_TYPE_STUDENT)
	unverifiedUser := user.NewUser(user.USER_TYPE_TEACHER)
	bannedUser := user.NewBannedUser(user.USER_TYPE_ADMIN)

	userRepo := user.NewMockRepository()
	userRepo.AddUsers([]*user.User{&validUser, &unverifiedUser, &bannedUser})
	jwtManager := auth.NewJwtManager(config.JWT{SecretKey: "secret_key", TTL: time.Minute})
	authService := auth.NewService(userRepo, jwtManager)
	authController := NewAuthController(authService)
	// SignIn Test Cases:
	testCases := []struct {
		name string
		form SingInRequest
		want int
	}{
		{
			name: "valid user (verified and permitted)",
			form: SingInRequest{validUser.Username, "secret"},
			want: http.StatusOK,
		},
		{
			name: "valid user with empty username and password",
			form: SingInRequest{"", ""},
			want: http.StatusBadRequest,
		},
		{
			name: "valid user with incorrect password",
			form: SingInRequest{validUser.Username, "invalid"},
			want: http.StatusUnauthorized,
		},
		{
			name: "valid user with incorrect username",
			form: SingInRequest{"invalid", "secret"},
			want: http.StatusUnauthorized,
		},
		{
			name: "unverified user with correct username and password",
			form: SingInRequest{unverifiedUser.Username, "secret"},
			want: http.StatusUnauthorized,
		},
		{
			name: "banned user with correct username and password",
			form: SingInRequest{bannedUser.Username, "secret"},
			want: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare context
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			jsonForm, _ := json.Marshal(tc.form)
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonForm))

			authController.SignIn(ctx)
			assert.Equal(t, tc.want, w.Code)

			if w.Code == http.StatusOK {
				var got map[string]SignInResponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatal(err)
				}
				data, dataExist := got["data"]
				if !dataExist {
					t.Fatal("key \"data\" not exist in json response")
				}
				assert.NotEmpty(t, data.AccessToken)
			}
		})
	}
}
