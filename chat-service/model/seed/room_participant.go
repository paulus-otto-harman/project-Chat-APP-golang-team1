package seed

import "project/chat-service/model"

func RoomParticipantSeed() []model.RoomParticipant {
	return []model.RoomParticipant{
		{RoomID: 1, UserID: 1},
		{RoomID: 1, UserID: 2},
		{RoomID: 2, UserID: 1},
		{RoomID: 3, UserID: 3},
		{RoomID: 4, UserID: 4},
	}
}
