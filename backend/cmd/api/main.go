package main

import (
	"flowgraph/internal/db"
	"flowgraph/internal/service"
)

type application struct {
	service *service.Service
}

func main() {
	// Initialize MongoDB database client (Panics if connection fails)
	database := db.InitMongoDbClient()
	// Initialize application
	app := &application{
		service: &service.Service{
			Db: database,
		},
	}

	// Start HTTP server
	ListenAndServe(app)
}
