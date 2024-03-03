package model

// Chat represents a chat room with a unique ID and name.
type Chat struct {
	ChatID   int64  // Unique identifier for the chat room.
	ChatName string // Name of the chat room.
}

// User represents a user participating in the chat system.
type User struct {
	UserID   int64  // Unique identifier for the user.
	UserName string // Name of the user.
	ChatID   int64  // ID of the chat room the user belongs to.
}

// Message represents a message sent within a chat room.
type Message struct {
	ChatID      int64  // ID of the chat room where the message is sent.
	UserName    string // Name of the user who sent the message.
	MessageText string // Text content of the message.
}

// ChatDelete represents a request to delete a chat room.
type ChatDelete struct {
	ChatID int64 // ID of the chat room to be deleted.
}
