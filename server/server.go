package main

import (
	"context"
	pb "grpc/pkg/telegram"
	"net"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"

	"google.golang.org/grpc"
)

var Bot *tgbotapi.BotAPI

type TelegramServer struct {
	pb.UnimplementedTelegramServer
}

func (s *TelegramServer) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {

	chatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	message := ""
	switch req.Level {
	case 0:
		message = ""
	case 1:
		message = "ALARM: "
	case 2:
		message = "WARNING: "
	case 3:
		message = "DEBUG: "
	case 4:
		message = "INFO: "
	}

	message = message + " " + req.GetMessage()

	msg := tgbotapi.NewMessage(chatID, message)

	_, err := Bot.Send(msg)

	if err == nil {
		log.Infof("Message successfully delivered")
	} else if err != nil {
		log.Errorf("Message delivery failed with error: %s", err)
	}
	return &pb.SendResponse{Message: req.GetMessage()}, nil
}

func main() {
	log.Info("Server is starting...")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	Bot = InitBot()
	pb.RegisterTelegramServer(grpcServer, &TelegramServer{})

	err = grpcServer.Serve(lis)
	if err != err {
		log.Fatal(err)
	}

}

func InitBot() *tgbotapi.BotAPI {
	token := os.Getenv("BOT_TOKEN")

	debug := os.Getenv("DEBUG")

	Bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic("Bot:", err)
	}

	Bot.Debug, _ = strconv.ParseBool(debug)

	return Bot
}
