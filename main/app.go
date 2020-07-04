package main

import (
	"enigma_mart_gerin/config/database"
	"enigma_mart_gerin/config/router"
	"fmt"
)

func main() {
	db := database.ConnectDB()
	r := router.CreateRouter()

	fmt.Println(r)
	fmt.Println(db)
	router.NewAppRouter(db, r).InitRouter()
	router.StartServer(r)
}
