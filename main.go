package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/routers"
	"github.com/shyptr/archiveofourown/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalln(err)
	}
	err = global.SetupDBEngine()
	if err != nil {
		log.Fatalln(err)
	}
	logger.SetupLogger()
	err = global.SetupTracer()
	if err != nil {
		log.Fatalln(err)
	}
}

// @title 同人圈
// @version 1.0
// @description 属于国内的同人圈
// @termsOfService https://github.com/shyptr/archiveofourown
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	server := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeOut * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("server.ListenAndServe error: %v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exit.")
}
