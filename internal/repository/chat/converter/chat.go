package converter

import (
	"chat-server/internal/model"
	modelRepo "chat-server/internal/repository/chat/model"
)

// FromRepoToChat converts a repository Chat model to a chat system Chat entity.
func FromRepoToChat(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ChatID:   chat.ChatID,
		ChatName: chat.ChatName,
	}
}

// FromChatToRepo converts a chat system Chat entity to a repository Chat model.
func FromChatToRepo(chat *model.Chat) *modelRepo.Chat {
	return &modelRepo.Chat{
		ChatID:   chat.ChatID,
		ChatName: chat.ChatName,
	}
}

// FromRepoToUser converts a repository User model to a chat system User entity.
func FromRepoToUser(u modelRepo.User) *model.User {
	return &model.User{
		UserID:   u.UserID,
		UserName: u.UserName,
		ChatID:   u.ChatID,
	}
}

// FromRepoToMessage converts a repository Message model to a chat system Message entity.
func FromRepoToMessage(m modelRepo.Message) *model.Message {
	return &model.Message{
		ChatID:      m.ChatID,
		UserName:    m.UserName,
		MessageText: m.MessageText,
	}
}
