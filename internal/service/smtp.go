package service

import (
	"fmt"
	"github.com/leshachaplin/emailSender/protocol"
	"net/smtp"
	"strconv"
)

type SMTPEmail struct {
	from     string
	to       string
	username string
	password string
	host     string
	port     int32
}

func NewSMTPEmail(port int32, from, to, username, password, host string) *SMTPEmail {
	return &SMTPEmail{
		from:     from,
		to:       to,
		username: username,
		password: password,
		host:     host,
		port:     port,
	}
}

//service.gmail.com

func (e *SMTPEmail) Send(data *protocol.UuidTemplate) error {
	//d, ok := data.(protocol.UuidTemplate)
	//if !ok {
	//	return fmt.Errorf("unable to convert data")
	//}
	err := smtp.SendMail(e.host+":"+strconv.Itoa(int(e.port)),
		smtp.PlainAuth("", e.username, e.password, e.host),
		e.from, []string{e.to}, []byte(fmt.Sprintf("<a href='%s'>click here</a>", data.Token)))
	return err
}
