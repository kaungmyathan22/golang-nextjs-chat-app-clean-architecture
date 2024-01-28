package main

import (
	"log"

	"github.com/kaungmyathan22/golang-nextjs-chat-app-clean-architecture/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	// userRep := user.NewRepository(dbConn.GetDB())
	// userSvc := user.NewService(userRep)
	// userHandler := user.NewHandler(userSvc)

	// hub := ws.NewHub()
	// wsHandler := ws.NewHandler(hub)
	// go hub.Run()

	// router.InitRouter(userHandler, wsHandler)
	// router.Start("0.0.0.0:8080")
}
