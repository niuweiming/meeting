package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"huiyi/models"
)

func SendMail(conf models.MailboxConf) error {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Sender)
	m.SetHeader("To", conf.RecipientList...)
	m.SetHeader("Subject", conf.Title)
	m.SetBody("text/html", conf.Body)
	m.SetHeader("Reply-To", conf.Sender)

	d := gomail.NewDialer(conf.SMTPAddr, conf.SMTPPort, conf.Sender, conf.SPassword)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 如果不验证证书，可以加上此行

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Error("发送邮件失败: ", err)
		return err
	}
	log.Info("邮件发送成功: ")
	return nil
}
