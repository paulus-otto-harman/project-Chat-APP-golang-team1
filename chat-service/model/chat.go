package model

import "gorm.io/gorm"

// Tabel Room
type Room struct {
	gorm.Model
	Name         string            `json:"name" gorm:"default:pv"`
	Participants []RoomParticipant `json:"participants" gorm:"foreignKey:RoomID"`
	Messages     []Message         `json:"messages" gorm:"foreignKey:RoomID"`
}

// Tabel RoomParticipant
type RoomParticipant struct {
	gorm.Model
	RoomID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Room   Room `gorm:"foreignKey:RoomID"` // Relasi ke Room
	User   User `gorm:"foreignKey:UserID"` // Relasi ke User
}

// Tabel User
type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"unique;not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Messages []Message `json:"messages" gorm:"foreignKey:SenderID"`
}

// Tabel Message
type Message struct {
	gorm.Model
	RoomID        uint    `json:"room_id"`
	SenderID      uint    `json:"sender_id"`
	Content       string  `json:"content"`
	AttachmentURL *string `json:"attachment_url"`
	ReplyTo       *uint   `json:"reply_to"`
	ReadAt        *string `json:"read_at"`
	Room          Room    `gorm:"foreignKey:RoomID"`   // Relasi ke Room
	Sender        User    `gorm:"foreignKey:SenderID"` // Relasi ke User
}

// Data Seeder untuk Room
func RoomSeed() []Room {
	return []Room{
		{Name: "General Room"},
		{Name: "Support Room"},
		{Name: "Private Room 1"},
		{Name: "Private Room 2"},
	}
}

// Data Seeder untuk User
func UserSeed() []User {
	return []User{
		{Username: "user1", Email: "user1@example.com", Password: "password123"},
		{Username: "user2", Email: "user2@example.com", Password: "password123"},
		{Username: "user3", Email: "user3@example.com", Password: "password123"},
		{Username: "user4", Email: "user4@example.com", Password: "password123"},
	}
}

// Data Seeder untuk RoomParticipant
func RoomParticipantSeed() []RoomParticipant {
	return []RoomParticipant{
		{RoomID: 1, UserID: 1},
		{RoomID: 1, UserID: 2},
		{RoomID: 2, UserID: 1},
		{RoomID: 3, UserID: 3},
		{RoomID: 4, UserID: 4},
	}
}

// Data Seeder untuk Message
func MessageSeed() []Message {
	return []Message{
		{RoomID: 1, SenderID: 1, Content: "Welcome to the General Room!"},
		{RoomID: 1, SenderID: 2, Content: "Hello, how can I assist you today?"},
		{RoomID: 2, SenderID: 1, Content: "I need help with my account."},
		{RoomID: 3, SenderID: 3, Content: "Hi, this is a private message."},
		{RoomID: 4, SenderID: 4, Content: "Private Room 2, message from User 4."},
	}
}
