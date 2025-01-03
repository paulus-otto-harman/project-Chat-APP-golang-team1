package main

import (
	"log"
	"net"
	"project/chat-service/handler"
	"project/chat-service/infra"

	pb "project/chat-service/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	// router := gin.Default()

	// Routes
	// router.GET("/ws", ctx.Ctl.ChatHandler.HandleWebSocket)
	// router.GET("/chat/history", ctx.Ctl.ChatHandler.GetChatHistory)
	// router.POST("/chat/read", ctx.Ctl.ChatHandler.MarkMessageAsRead)

	// Start server
	ctx.Log.Info("starting server on :8080")
	// if err := router.Run(":8080"); err != nil {
	// 	ctx.Log.Fatal("server failed to start", zap.Error(err))
	// }
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &handler.ChatHandler{
		Service: ctx.Ctl.ChatHandler.Service,
		Logger:  ctx.Log,
	})

	if err := grpcServer.Serve(lis); err != nil {
		return
	}
}
