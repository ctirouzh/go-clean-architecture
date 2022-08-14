package controller

import (
	"lms/internal/app/auth"
	"lms/internal/domain/user"
	"lms/internal/pkg/apperr"
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
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}

	usr, err := ctrl.authService.SignUp(req.Username, req.Email, req.Password)
	if err != nil {
		status := http.StatusInternalServerError
		if err == user.ErrEmailAlreadyTaken || err == user.ErrUsernameAlreadyTaken {
			status = http.StatusConflict
		}
		c.Error(apperr.New(status, err.Error()))
		return
	}

	var res *UserDTO
	res.Prepare(usr)
	c.JSON(http.StatusOK, gin.H{"data": res})
}
