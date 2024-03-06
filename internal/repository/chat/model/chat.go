package repomodel

import "time"

// Chat represents a chat room with a unique ID and name in repo layer
type Chat struct {
	ChatName string
}

// User represents a user participating in the chat system for repo layer
type User struct {
	UserID   int64
	UserName string
	ChatID   int64
}

// Message represents a message sent within a chat room for repo layer
type Message struct {
	MessageID   int64
	ChatID      int64
	UserName    string
	MessageText string
	CreatedAt   time.Time
}
