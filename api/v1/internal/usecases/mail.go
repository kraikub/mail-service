package usecases

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/smtp"
	"text/template"
)

type MailUseCase interface {
	TwoFA(lang string, to string, code string, name string, deviceName string) error
	VerifyEmail(lang string, to string, code string, name string) error
}

type MailData struct {
	To string
}

type TwoFactorMailStructure struct {
	To   string
	Code string
	Name string
	DeviceName string
}

type mailUseCase struct {
	smtpHost             string
	smtpPort             string
	serviceEmail         string
	serviceEmailPassword string
}

func CreateMailUseCase(smtpHost string, smtpPort string, serviceEmail string, serviceEmailPassword string) mailUseCase {
	return mailUseCase{
		smtpHost:             smtpHost,
		smtpPort:             smtpPort,
		serviceEmail:         serviceEmail,
		serviceEmailPassword: serviceEmailPassword,
	}
}

func loadAndParse(templatePath string, mdata interface{}) ([]byte, error) {
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return []byte{}, err
	}
	t, err := template.New("mail").Parse(string(content))

	if err != nil {
		return []byte{}, err
	}
	
	buf := new(bytes.Buffer)
	t.Execute(buf, mdata)
	return buf.Bytes(), nil
}

func (s mailUseCase) TwoFA(lang string, to string, code string, name string, deviceName string) error {

	if lang == "" {
		return errors.New("Cannot find language prefference.")
	}

	mdata := TwoFactorMailStructure{
		To: to,
		Code: code,
		Name: name,
		DeviceName: deviceName,
	}

	from := s.serviceEmail
	password := s.serviceEmailPassword

	path := "./templates/"+lang+"/2fa.txt"

	message, err := loadAndParse(path, mdata)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Create authentication
	auth := smtp.PlainAuth("", from, password, s.smtpHost)
	// Send actual message
	go func() {
		err = smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, from, []string{to}, message)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (s mailUseCase) VerifyEmail(lang string, to string, code string, name string) error {

	if lang == "" {
		return errors.New("Cannot find language prefference.")
	}

	mdata := TwoFactorMailStructure{
		To: to,
		Code: code,
		Name: name,
	}

	from := s.serviceEmail
	password := s.serviceEmailPassword

	path := "./templates/"+lang+"/verify.txt"

	message, err := loadAndParse(path, mdata)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Create authentication
	auth := smtp.PlainAuth("", from, password, s.smtpHost)
	// Send actual message
	go func() {
		err = smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, from, []string{to}, message)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}
