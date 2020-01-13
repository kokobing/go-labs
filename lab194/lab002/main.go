package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"mime"
	"net/mail"
	"net/smtp"
)

func encodeRFC2047(s string) string {
	return mime.QEncoding.Encode("utf-8", s)
}

func main() {
	// Set up authentication information.

	smtpServer := "smtp.qq.com"
	auth := smtp.PlainAuth(
		"",
		"用户名",
		"密码",
		smtpServer,
	)

	from := mail.Address{"IFTTT", "123@qq.com"}
	to := mail.Address{"收件人", "123@mail.dida365.com"}
	title := "当前时段统计报表"

	body := "报表内容一切正常"
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
