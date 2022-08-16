package http

import (
	"lms/internal/adapters/controller"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(r *gin.Engine, ctrl *controller.Auth) {
	auth := r.Group("/api/auth")
	auth.POST("/signup", ctrl.SignUp)
}
