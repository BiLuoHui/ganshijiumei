package main

import (
	"context"
	"database/sql"
	"github.com/BiLuoHui/ganshijiumei/ent"
	ents "github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	dsn             = "host=127.0.0.1 port=5432 user=hui password= dbname=hui sslmode=disable"
	maxIdleConns    = 6
	maxOpenConns    = 100
	connMaxLifetime = time.Hour * 2
)

var client *ent.Client

func main() {
	migrate(context.Background())

	defer client.Close()
}

func migrate(ctx context.Context) {
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v\n", err)
	}
}

func init() {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	drv := ents.OpenDB("postgres", db)

	client = ent.NewClient(ent.Driver(drv))
}
