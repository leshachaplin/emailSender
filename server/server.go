package server

import (
	"context"
	"github.com/leshachaplin/emailSender/protocol"
	"github.com/leshachaplin/emailSender/service"
	log "github.com/sirupsen/logrus"
)

type Server struct {
}

func (s *Server) Send(ctx context.Context, req *protocol.SendRequest) (*protocol.EmptyResponse, error)  {
	smtpReq := service.NewSMTPEmail(req.Email.Port, req.Email.From, req.Email.To,
		req.Email.Username, req.Email.Password, req.Email.Host)

	err := smtpReq.Send(req.Template)
	if err != nil {
		log.Errorf("error in sending email: %s", err)
		return nil, err
	}
	return &protocol.EmptyResponse{}, nil
}
