package main

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"strings"
)

func main() {
	// 邮件信息
	from := "auth@3dxr.com" // 发件人邮箱
	username := "auth@3dxr.com"
	password := "zQp8Jf8S7CxMQ1nT"                       // 发件人邮箱密码，有些邮箱提供的是授权码，而不是真正的密码
	to := []string{"237230999@qq.com", "OOXX666@88.com"} // 收件人邮箱
	subject := "Test Email"                              // 邮件主题
	body := "This is the email body."                    // 邮件内容

	// SMTP服务器地址和端口
	smtpHost := "smtp.feishu.cn"
	smtpPort := "465"

	// 组装邮件内容
	message := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// 建立安全连接
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpHost,
	}
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 执行SMTP操作
	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	// 构建认证信息
	auth := smtp.PlainAuth("", username, password, smtpHost)
	if err = c.Auth(auth); err != nil {
		log.Fatal(err)
	}

	if err = c.Mail(from); err != nil {
		log.Fatal(err)
	}

	for _, recipient := range to {
		err = c.Rcpt(strings.TrimSpace(recipient))
		if err != nil {
			log.Fatal(err)
		}
	}

	w, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent successfully!")
}
