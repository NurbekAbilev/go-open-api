package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	sw "github.com/flowchartsman/swaggerui"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/config"
	"github.com/nurbekabilev/go-open-api/internal/db"
	"github.com/nurbekabilev/go-open-api/internal/handler"

	"github.com/nurbekabilev/go-open-api/internal/handler/swaggerui"
)

func main() {
	config.LoadDotEnv()
	db.Migrate(os.Getenv("DB_URL"))

	cfg, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.InitPgxConnect(context.Background(), cfg)
	if err != nil {
		log.Fatal("could not init connect %w", err)
	}

	err = app.InitApp(app.AppConfig{PgxCon: conn})
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handler.GetRoutes())
	http.Handle("/api/swagger/", http.StripPrefix("/api/swagger", sw.Handler(swaggerui.GetSwaggerYml())))

	fmt.Println("Listening to the:", config.Host)
	err = http.ListenAndServe(config.Host, nil)
	if err != nil {
		panic(err)
	}
}
