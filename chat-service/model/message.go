package model

import "gorm.io/gorm"

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
