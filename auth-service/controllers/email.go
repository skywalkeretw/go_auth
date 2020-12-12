package controllers

/*
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("models/emailTemplates/*.gohtml"))
}

func SendConfirmationEmail(email string) () {

	msg := GetEmailContent("confirmationEmailTemplate.gohtml", email)
	err := smtp.SendMail("", AuthMail(), "", strings.Split(email, ""), msg  )
	if err != nil {
		log.Println("Email Error:", err)
	}
}

func SendPasswordResetMail(email string)  {
	msg := GetEmailContent("resetPasswordEmailTemplate.gohtml", email)
	err := smtp.SendMail("", AuthMail(), "", strings.Split(email, ""), msg  )
	if err != nil {
		log.Println("Email Error:", err)
	}
}


func AuthMail() smtp.Auth {
	return smtp.PlainAuth("", "","", "")
}

func GetEmailContent(tmpFile string, data interface{}) []byte {
	t := template.New("action")

	var err error
	t, err = t.ParseFiles("models/" + tmpFile )
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
*/