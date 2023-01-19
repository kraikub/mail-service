package controllers

type MailTransferDTO struct {
	To          string `json:"to"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	DeviceName  string `json:"deviceName"`
	Lang        string `json:"lang"`
	Ref         string `json:"ref"`
	OrgName     string `json:"orgName"`
	OrgUsername string `json:"orgUsername"`
	By          string `json:"by"`
	Position string `json:"position"`
}
