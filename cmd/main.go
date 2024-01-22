package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/rogudator/tuvebes-backend/internal/repository"
	"github.com/rogudator/tuvebes-backend/internal/transport"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	config := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	db, err := repository.NewPosgresDB(config)
	if err != nil {
		log.Fatal(err)
	}
	r := repository.NewRepository(db)
	handler := transport.NewTransport(r)
	router := handler.InitRoutes()

	router.Run("0.0.0.0:8081")
	server := &http.Server{
		Addr:         "0.0.0.0" + ":" + "8081",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("error occured when runnting http server: %s", err.Error())
		}
	}()

	log.Println("tuvebes_backend started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("tuvebes_backend shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db disconnection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
