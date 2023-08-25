package otptimize

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

// mail verification
func sendVerificationEmail(appName string, targetName string, targetEmail string, token string) {
	data := map[string]string{
		"Name":    targetName,
		"AppName": appName,
		"TOKEN":   token,
	}

	// get template file
	t, err := template.ParseFiles("templates/verificationEmail.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bodyMail := new(bytes.Buffer)

	// executing template, and parse "data" to template
	t.Execute(bodyMail, data)

	// create new message
	verificationEmail := gomail.NewMessage()
	verificationEmail.SetHeader("From", appName)
	verificationEmail.SetHeader("To", targetEmail)
	verificationEmail.SetHeader("Subject", "Email Verification")
	verificationEmail.SetBody("text/html", bodyMail.String())

	err = MailConnection.DialAndSend(verificationEmail)
	if err != nil {
		fmt.Println("Failed to send verification email")
		fmt.Println(err.Error())
		return
	}

	log.Println("Email sent successfully !")
}
