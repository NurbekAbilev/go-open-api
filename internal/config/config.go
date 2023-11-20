package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nurbekabilev/go-open-api/internal/fs"
)

const Host = ":8080"

func LoadDotEnv() {
	rootPath := fs.RootPath()
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}
