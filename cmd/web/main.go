package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	sw "github.com/flowchartsman/swaggerui"
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
	closer, err := app.InitApp(app.AppConfig{})
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	http.Handle("/", handler.GetRoutes())
	http.Handle("/api/swagger/", http.StripPrefix("/api/swagger", sw.Handler(swaggerui.GetSwaggerYml())))

	fmt.Println("Listening to the:", config.Host)
	err = http.ListenAndServe(config.Host, nil)
	if err != nil {
		panic(err)
	}
}
