package converter

import (
	"github.com/mchekalov/chat-server/internal/model"
	desc "github.com/mchekalov/chat-server/pkg/chat_api_v1"
)

// CreateChatInput represents the input structure for creating a chat room.
type CreateChatInput struct {
	ChatID   uint64
	ChatName string
}

// CreateChatOutput represents the output structure for creating a chat room.
type CreateChatOutput struct {
	ID int64
}

// SendMessageInput represents the input structure for sending a message.
type SendMessageInput struct {
	From string
	Text string
}

// ToCreateChatInput converts a CreateRequest object from the API to a model.Chat entity.
func ToCreateChatInput(req *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		ChatName: req.Chatname,
	}
}

// FromCreateChatInput converts an ID from model.Chat to a CreateResponse object for the API.
func FromCreateChatInput(input int64) *desc.CreateResponse {
	return &desc.CreateResponse{
		Id: input,
	}
}

// ToDeleteChatInput converts a DeleteRequest object from the API to a model.ChatDelete entity.
func ToDeleteChatInput(req *desc.DeleteRequest) *model.ChatDelete {
	return &model.ChatDelete{
		ChatID: req.GetId(),
	}
}

// ToSendMessageInput converts a SendMessageRequest object from the API to a model.Message entity.
func ToSendMessageInput(req *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		ChatID:      0,
		UserName:    req.Info.From,
		MessageText: req.Info.Text,
	}
}
