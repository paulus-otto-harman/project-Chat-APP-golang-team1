package seed

import "project/chat-service/model"

func MessageSeed() []model.Message {
	return []model.Message{
		{RoomID: 1, SenderID: 1, Content: "Welcome to the General Room!"},
		{RoomID: 1, SenderID: 2, Content: "Hello, how can I assist you today?"},
		{RoomID: 2, SenderID: 1, Content: "I need help with my account."},
		{RoomID: 3, SenderID: 3, Content: "Hi, this is a private message."},
		{RoomID: 4, SenderID: 4, Content: "Private Room 2, message from User 4."},
	}
}
