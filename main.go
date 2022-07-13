package main

import (
	"context"
	"hermes-mail-svc/db"
	"hermes-mail-svc/mailer"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sc
		cancel()
	}()

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		os.Exit(1)
	}

	app := App{}
	app.Initialize(
		ctx,
		db.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     int32(dbPort),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Catalog:  os.Getenv("DB_NAME"),
		},
		mailer.Config{
			Sender:          os.Getenv("MAIL_SENDER"),
			Region:          os.Getenv("MAIL_REGION"),
			AccessKeyID:     os.Getenv("MAIL_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("MAIL_SECRET_ACCESS_KEY"),
		},
	)

	app.Run()

	<-ctx.Done()
}
