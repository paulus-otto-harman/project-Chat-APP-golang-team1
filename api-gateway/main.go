package main

import (
	"log"
	"project/api-gateway/infra"
	"project/api-gateway/routes"
	_ "project/docs"
)

// @title Project Microservice App Chat System
// @version 1.0
// @description
// @termsOfService http://example.com/terms/
// @contact.name Team 1
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email lumoshive.academy@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description <h5>Please type "Bearer", following with a space, and then the JWT token</h5>
func main() {
	var err error

	var ctx *infra.ServiceContext
	if ctx, err = infra.NewServiceContext(); err != nil {
		log.Fatal("can't init service context %w", err)
	}

	routes.NewRoutes(*ctx)
}
