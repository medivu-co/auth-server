package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"medivu.co/auth/envs"
	"medivu.co/auth/logger"
)

var Conn *pgx.Conn

func Connect() {
	var err error

	// Connect to the database
	Conn, err = pgx.Connect(context.Background(), envs.PostgresDBURL())

	if err != nil {
		logger.Get().Fatal(err.Error())
	}

	// Check if the connection is successful
	if err = Conn.Ping(context.Background()); err != nil {
		logger.Get().Fatal(err.Error())
	}

	// Set connection pool settings
	// DB.SetMaxOpenConns(10)   // Set the maximum number of open connections
	// DB.SetMaxIdleConns(5)    // Set the maximum number of idle connections
	// DB.SetConnMaxLifetime(0) // Set the maximum connection lifetime (0 means unlimited)

	logger.Get().Info("Connected to database")
}

func Close() {
	if err := Conn.Close(context.Background()); err != nil {
		logger.Get().Fatal(err.Error())
	}
}
