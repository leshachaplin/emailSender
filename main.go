package main

import (
	"context"
	"fmt"
	"github.com/leshachaplin/emailSender/internal/config"
	"github.com/leshachaplin/emailSender/internal/server"
	"github.com/leshachaplin/emailSender/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(cfg.GrpcPort))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", cfg.GrpcPort)

	_, cnsl := context.WithCancel(context.Background())

	s := grpc.NewServer()
	srv := &server.Server{}
	emailService := protocol.EmailServiceService{
		Send: srv.Send,
	}
	protocol.RegisterEmailServiceService(s, &emailService)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("server is not connect %s", err)
		}
	}()

	for {
		select {
		case <-c:
			cnsl()
			log.Info("Cansel is succesful")
			close(c)
			return
		}
	}


}
