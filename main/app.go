package main

import (
	"enigma_mart_gerin/config/database"
	"enigma_mart_gerin/config/router"
)

func main() {
	db := database.ConnectDB()
	r := router.CreateRouter()

	router.NewAppRouter(db, r).InitRouter()
	router.StartServer(r)
}
