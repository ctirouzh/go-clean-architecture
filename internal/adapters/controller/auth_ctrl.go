package controller

import (
	"errors"
	"lms/internal/app/auth"
	"lms/internal/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	authService *auth.Service
}

func NewAuthController(authService *auth.Service) *Auth {
	return &Auth{authService: authService}
}

func (ctrl *Auth) SignUp(c *gin.Context) {

	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	usr, err := ctrl.authService.SignUp(req.Username, req.Email, req.Password)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, user.ErrUsernameOrEmailAlreadyTaken) {
			status = http.StatusConflict
		}
		c.AbortWithError(status, err)
		return
	}

	var res UserDTO
	res.Prepare(*usr)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (ctrl *Auth) SignIn(c *gin.Context) {
	var req SingInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := ctrl.authService.SignIn(req.Username, req.Password)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": SignInResponse{AccessToken: token}})
}
