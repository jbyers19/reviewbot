package reviewbot

import (
	"context"
	"fmt"

	"github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Message is a gRPC server that sends messages to customers.
// It implements the pb.MessageServer interface.
type Message struct {
	pb.UnimplementedMessageServer
}

// Send sends a custom message generated from a message template to a customer.
func (m *Message) Send(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	_, ok := TemplatesDB.TemplatesMap[in.TemplateName]
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("template %s does not exist", in.TemplateName))
	}

	msg, err := TemplatesDB.GenerateMessage(in.TemplateName, in.FirstName, in.LastName)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error generating message: %s", err.Error()))
	}

	err = SendTelegramMessage(in.FirstName, in.LastName, msg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.SendMessageResponse{Message: msg}, nil
}
