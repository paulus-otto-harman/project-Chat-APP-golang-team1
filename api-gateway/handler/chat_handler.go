package handler

import (
	"context"
	"log"
	"net/http"
	"project/api-gateway/database"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	"project/api-gateway/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ChatController struct {
	service service.Service
	logger  *zap.Logger
	rdb     database.Cacher
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins. Modify as per your security needs.
		return true
	},
}

func NewChatController(service service.Service, logger *zap.Logger, rdb database.Cacher) *ChatController {
	return &ChatController{service, logger, rdb}
}

func (ctrl *ChatController) Websocket(c *gin.Context) {
	username := c.GetString("email")
	if username == "" {
		BadResponse(c, "unauthorized", http.StatusUnauthorized)
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		BadResponse(c, "webSocket upgrade failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	roomId := c.Param("id")
	pubsub := ctrl.rdb.Subcribe("room:" + roomId)
	defer pubsub.Close()

	for {
		payload, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println("failed to received message from redis: ", err.Error())
			return
		}

		if err = conn.WriteMessage(websocket.TextMessage, []byte(payload.Payload)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}

}

func (ctrl *ChatController) NewMessage(c *gin.Context) {
	message := model.Message{Sender: c.GetString("email")}
	if message.Sender == "" {
		BadResponse(c, "unauthorized", http.StatusUnauthorized)
		return
	}
	roomId := c.Param("id")
	uintRoomId, err := helper.Uint(roomId)
	if err != nil {
		BadResponse(c, "invalid chat id", http.StatusBadRequest)
		return
	}
	message.RoomId = uintRoomId

	if err = c.ShouldBindJSON(&message); err != nil {
		BadResponse(c, "invalid chat content", http.StatusUnprocessableEntity)
		return
	}

	if err = ctrl.service.Chat.SaveMessage(&message); err != nil {
		log.Println(err, "error from grpc save message")
		return
	}

	ctrl.rdb.Publish("room:"+roomId, message.Content)

	GoodResponseWithData(c, "message sent", http.StatusOK, message)
}

func (ctrl *ChatController) GetRoomMessages(c *gin.Context) {
	query := c.Query("page")
	var page uint
	var err error
	if query != "" {
		page, err = helper.Uint(query)
		if err != nil {
			BadResponse(c, err.Error(), http.StatusBadRequest)
			return
		}
	}
	param := c.Param("id")
	roomId, err := helper.Uint(param)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := ctrl.service.Chat.GetRoomMessages(roomId, page)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get History Message Succes", http.StatusOK, res)
}

func (ctrl *ChatController) GetAllParticipants(c *gin.Context) {
	param := c.Param("id")
	roomId, err := helper.Uint(param)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := ctrl.service.Chat.GetRoomParticipants(roomId)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get Participants Success", http.StatusOK, res)
}

func (ctrl *ChatController) AddParticipants(c *gin.Context) {
	param := c.Param("id")
	roomId, err := helper.Uint(param)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	var participant model.Participant
	if err := c.ShouldBindJSON(&participant); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := ctrl.service.Chat.AddRoomParticipant(uint64(roomId), participant.Email)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get Participants Success", http.StatusOK, res)
}
