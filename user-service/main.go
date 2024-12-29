package main

import (
	"log"
	"project/user-service/infra"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	log.Println(ctx)
}
