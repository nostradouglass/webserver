package main

import (
	"log"
	"net/http"
	"webserver/router"
	"webserver/services"
	"webserver/utils"
)

func main() {

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", appRouter))
}
