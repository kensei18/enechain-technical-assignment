package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kensei18/enechain-technical-assignment/app/graph/admin"
	"github.com/kensei18/enechain-technical-assignment/app/graph/admin/resolver"
	"github.com/kensei18/enechain-technical-assignment/app/loggers"
	"github.com/kensei18/enechain-technical-assignment/app/server"
	"github.com/kensei18/enechain-technical-assignment/app/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8081"

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	logger := loggers.NewDefaultLogger(os.Stdout, slog.LevelDebug)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open(
		postgres.Open("dbname=app host=localhost port=5432 user=postgres password=password sslmode=disable"),
		&gorm.Config{Logger: loggers.NewGormLogger(logger)},
	)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			slog.Error(fmt.Sprintf("failed to close database connection: %v\n", err))
		}
		if err = sqlDB.Close(); err != nil {
			slog.Error(fmt.Sprintf("failed to close database connection: %v\n", err))
		}
		slog.Info("database connection was closed")
	}()
	dbFunc := func(ctx context.Context) *gorm.DB { return db.WithContext(ctx) }

	s := &server.GraphQLServer{
		Port: port,
		Schema: admin.NewExecutableSchema(
			admin.Config{Resolvers: &resolver.Resolver{DB: dbFunc}},
		),
		Logger:  loggers.NewDefaultLogger(os.Stdout, slog.LevelDebug),
		Loaders: storage.NewLoaders(&storage.Reader{DB: dbFunc}),
	}

	go func() { s.Serve() }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
