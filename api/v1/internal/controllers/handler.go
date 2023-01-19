package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kraikub/mail-service/api/v1/internal/usecases"
)

type handler struct {
	mail usecases.MailUseCase
}

func (handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func (h handler) TwoFA(c *gin.Context) {
	var mDTO MailTransferDTO

	if err := c.ShouldBind(&mDTO); err != nil {
		handleResponse(c, http.StatusUnprocessableEntity, false, err.Error(), nil)
		return
	}
	h.mail.TwoFA(mDTO.Lang, mDTO.To, mDTO.Code, mDTO.Name, mDTO.DeviceName, mDTO.Ref)
	c.JSON(http.StatusOK, gin.H{
		"email_status": "sent",
	})
}

func (h handler) VerifyEmail(c *gin.Context) {
	var mDTO MailTransferDTO

	if err := c.ShouldBind(&mDTO); err != nil {
		handleResponse(c, http.StatusUnprocessableEntity, false, err.Error(), nil)
		return
	}
	h.mail.VerifyEmail(mDTO.Lang, mDTO.To, mDTO.Code, mDTO.Name)
	c.JSON(http.StatusOK, gin.H{
		"email_status": "sent",
	})
}

func (h handler) OrgInvite(c *gin.Context) {
	var mDTO MailTransferDTO

	if err := c.ShouldBind(&mDTO); err != nil {
		handleResponse(c, http.StatusUnprocessableEntity, false, err.Error(), nil)
		return
	}
	h.mail.OrgInvite(mDTO.To, mDTO.Code, mDTO.Name, mDTO.OrgName, mDTO.OrgUsername, mDTO.By, mDTO.Position)
	c.JSON(http.StatusOK, gin.H{
		"email_status": "sent",
	})
}
