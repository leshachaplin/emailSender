package server

import (
	"context"
	"fmt"
	conf "github.com/leshachaplin/config"
	"github.com/leshachaplin/emailSender/protocol"
	"google.golang.org/grpc"
	"testing"
)

func TestSendEmail(t *testing.T) {
	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial("0.0.0.0:50052", opts)
	if err != nil {
		t.Error(err)
	}
	defer clientConnInterface.Close()
	client := protocol.NewEmailServiceClient(clientConnInterface)

	awsConf, err := conf.NewForAws("us-west-2")
	if err != nil {
		t.Fatal("Can't connect to aws", err)
	}

	smtp, err := awsConf.GetSMTP("aws-ssm-smtp://SMTP/")
	if err != nil {
		t.Fatal("Can't get smtp connection string", err)
	}

	email := &protocol.SMTPEmail{
		From:     smtp.Email,
		To:       smtp.Username,
		Username: smtp.Username,
		Password: smtp.Password,
		Host:     smtp.Host,
		Port:     int32(smtp.Port),
	}

	template := &protocol.UuidTemplate{Token: "hello"}

	requestSendMail := &protocol.SendRequest{
		Email:    email,
		Template: template,
	}

	_, err = client.Send(context.Background(), requestSendMail)
	if err == nil {
		fmt.Printf("EMAIL SEND")
	} else {
		t.Errorf("ERROR: %s ", err)
	}
}
