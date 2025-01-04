package main

import (
	"api_gateway/infra"
	"api_gateway/routes"
	"log"
)

func main() {
	var err error

	var ctx *infra.ServiceContext
	if ctx, err = infra.NewServiceContext(); err != nil {
		log.Fatal("can't init service context %w", err)
	}

	routes.NewRoutes(*ctx)
}
