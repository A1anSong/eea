package main

import (
	"context"
	"eea/config"
	"eea/model"
	"eea/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./")
	viper.SetConfigName("server")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = viper.Unmarshal(&config.Configs)
	if err != nil {
		log.Fatal(err.Error())
	}
	model.ConnectDB()

	r := router.InitRouter()

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
