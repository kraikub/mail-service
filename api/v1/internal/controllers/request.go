package controllers

type MailTransferDTO struct {
	To         string `json:"to"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	DeviceName string `json:"deviceName"`
	Lang       string `json:"lang"`
}
