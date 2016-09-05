package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/gatorloopwebapp/database"
	"github.com/gatorloopwebapp/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	server.RegisterAPI("/api", wsContainer)
	database.InitDB()

	log.Info("hosting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
