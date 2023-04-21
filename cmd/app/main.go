package main

import (
	"context"
	"control-accounting-service/internal/config"
	"control-accounting-service/internal/delivery/http"
	storage "control-accounting-service/internal/repository/storage/postgres"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"time"
)

func main() {
	log.Println("start")
	time.Sleep(2 * time.Second)
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.Dsn)))

	db := bun.NewDB(pgdb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	repo, _ := storage.New(db)

	server := http.New(ctx, gin.Default(), repo)
	server.InitRouter()

	err = server.Start(cfg)
	if err != nil {
		log.Println(err)
	}
}
