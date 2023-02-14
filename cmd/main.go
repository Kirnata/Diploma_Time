package main

import (
	"context"
	"github.com/Kirnata/Diploma_Time/internal/entity"
	"github.com/Kirnata/Diploma_Time/internal/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handlers := handler.NewHandler()
	srv := new(entity.Server)
	go func() {
		err := srv.Run("8080", handlers.InitRouts())
		if err != nil {
			log.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()
	log.Println("Server started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Server finished")

	err := srv.ShutDown(context.Background())
	if err != nil {
		log.Printf("error occurred on server shutting down: %s\n", err)
	}
}
