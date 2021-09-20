package main

import (
	"TechTestNoSql/infraestructure/drivenadapters"
	"TechTestNoSql/infraestructure/handlers"
	"net/http"
)

func init() {
	// Driven Adapter Redis
	dbVote := drivenadapters.BirdVoteRedis{}
	// Driven Adapter Rest to validate Bird existence on the firs mircoservice
	dbValidate := drivenadapters.BirdValidateExistsRest{}

	mainHandler := handlers.MainHandler{}
	mainHandler.Start(dbValidate, dbVote)

}


func main() {
	// Use default handler initialized on init
	http.ListenAndServe(":8087", nil)
}