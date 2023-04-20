package main

import (
	"context"
	"control-accounting-service/internal/delivery/http"
	"control-accounting-service/internal/delivery/http/config"
	storage "control-accounting-service/internal/repository/storage/postgres"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

func main() {
	ctx := context.Background()

	dsn := "postgres://postgres:postgres@localhost:5432/test?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	migrations("file://db/migrations", dsn)

	db := bun.NewDB(pgdb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	repo, _ := storage.New(db)

	server := http.New(ctx, gin.Default(), repo)
	server.InitRouter()
	err := server.Start(config.New("localhost", 8080))
	if err != nil {
		log.Println(err)
	}
}

func migrations(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}
	log.Println("database migrated successfully")
}
