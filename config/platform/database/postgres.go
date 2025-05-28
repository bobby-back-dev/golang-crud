package database

import (
	"context"
	"fmt"
	"github.com/bobby-back-dev/golang-crud/config/platform/godo"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

var Pool *pgxpool.Pool

func ConnectToDb() error {

	dbUrl := godo.GetEnv("DBURL")
	fmt.Println("DB URL:", dbUrl)
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return fmt.Errorf("could not parse postgres url: %v", err)
	}
	config.MaxConns = 10
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour
	config.HealthCheckPeriod = time.Minute

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("pgxpool.New: %w", err)
	}

	Pool = conn
	log.Println("Successfully connected to database")
	return nil
}

func GetPool() *pgxpool.Pool {
	if Pool == nil {
		err := ConnectToDb()
		if err != nil {
		}
	}
	return Pool
}

func ClosePool() {
	if Pool != nil {
		Pool.Close()
		log.Println("Successfully closed database")
	}
}
