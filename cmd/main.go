package main

import (
	"context"
	_ "github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"xpay"
	"xpay/configs"
	"xpay/pkg/handler"
	"xpay/pkg/repository"
	"xpay/pkg/service"
)

func main() {
	err := configs.Init()
	if err != nil {
		logrus.Fatalf("Error init config: %s", err)
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error load enviroment var: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Error database connection: %s", err)
	}

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handlers := handler.NewHandler(serv)

	s := new(xpay.Server)

	go func() {
		logrus.Fatal(s.Run(viper.GetString("port"), handlers.InitRoutes()))
	}()

	logrus.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server turned off")

	err = s.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("error on turning off server: %s", err)
	}

	err = db.Close()
	if err != nil {
		logrus.Errorf("error on db close: %s", err)
	}
}
