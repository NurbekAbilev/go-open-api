package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nurbekabilev/go-open-api/internal/fs"
)

func LoadDotEnv() {
	rootPath := fs.RootPath()
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal(err)
	}
}
