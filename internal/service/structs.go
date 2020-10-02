package service

type PasswordTemplate struct {
	Token string `json:"token"`
}

type EmailSender interface {
	Send(templateId interface{}) error
}
