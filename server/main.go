package main

import (
	"context"
	"eea/config"
	"eea/model"
	"eea/router"
	"eea/util"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {

	config.InitConfig()
	initLog()
	util.Init()
	model.InitDB()
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

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	lv, err := log.ParseLevel(config.Configs.Log.Level)
	if err != nil {
		log.Errorf("parse log level failed, use info level,err: %s", err.Error())
		return
	}
	log.SetLevel(lv)
	log.SetReportCaller(true)
	fileName := config.Configs.Log.File
	logF, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal("open log file %s error:", fileName, err.Error())
	}
	log.SetOutput(logF)
}
