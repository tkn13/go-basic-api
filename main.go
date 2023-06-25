package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThaksinCK/go-basic-api.git/config"
	"github.com/ThaksinCK/go-basic-api.git/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := handler.Routes{}
	handleRoute := r.InitRouter()
	Appsrv := &http.Server{Addr: config.Port, Handler: handleRoute}

	go func() {
		var err error = nil
		err = Appsrv.ListenAndServe()
		if err != nil {
			log.Fatalf("[Main] Unable to start server : %+v\n", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

}
