package controllers

import (
	"bytes"
	"fmt"
	"github.com/skywalkeretw/auth/auth"
	"github.com/skywalkeretw/auth/models"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var tpl *template.Template

type smtpServer struct {
	host string
	port string
	from string
	password string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
func (s *smtpServer)AuthMail() smtp.Auth {
	return smtp.PlainAuth(s.Address(), s.from,s.password, s.host)
}

func SendMail(to []string, subject string, message string)  {
	smtpServer := smtpServer{
		host:     "smtp.gmail.com",
		port:     "587",
		from:     os.Getenv("EMAIL_ADDRESS"),
		password: os.Getenv("EMAIL_PASSWORD"),
	}    // Message.

	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message))
	err := smtp.SendMail(smtpServer.Address(), smtpServer.AuthMail(), smtpServer.from, to, msg)
	if err != nil {
		log.Println("Email Error:", err)
		return
	}
	log.Println("Email Sent!")
}



func SendConfirmationEmail(user *models.User) {

	token, err := auth.CreateConfirmToken(*user)
	if err != nil {
		log.Println("Confirm Token Error:", err)
	}
	log.Println("confirm token:", token)
	/*msg := GetEmailContent("confirmationEmailTemplate.gohtml", struct {
		Token string
	}{Token: token})*/
	subject := "Account Verifivation Mail"
	message := "Hello " + strings.Title(user.Firstname) + " " + strings.Title(user.Lastname) + ",\n"+
		"Please follow the link to confirm your account.\n"+ "link: " + "localhost:8080/confirmUser?token="+token
	SendMail([]string{user.Email}, subject, message)
}
/*
func SendPasswordResetMail(email string)  {
	token, err := auth.CreateConfirmToken(*user)
	if err != nil {
		log.Println("Confirm Token Error:", err)
	}
	log.Println("confirm token:", token)

	subject := "Account Verifivation Mail"
	message := "Hello " + strings.Title(user.Firstname) + " " + strings.Title(user.Lastname) + ",\n"+
		"Please follow the link to confirm your account.\n"+ "link: " + "localhost:8080/confirmUser?token="+token
	SendMail([]string{user.Email}, subject, message)
}
*/



func GetEmailContent(tmpFile string, data interface{}) []byte {
	t := template.New("action")

	var err error
	t, err = t.ParseFiles("models/emailTemplates/" + tmpFile )
	if err != nil {
		// return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		// return err
	}

	result := tpl.String()
	return []byte(result)
}
