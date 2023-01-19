package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kraikub/mail-service/api/v1/internal/usecases"
)

func AssignRouter(r *gin.Engine, mail usecases.MailUseCase) {
	h := handler{
		mail: mail,
	}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", h.Hello)
		v1.POST("/2fa", h.TwoFA)
		v1.POST("/verify-email", h.VerifyEmail)
		v1.POST("/org-invite", h.OrgInvite)
	}

}
