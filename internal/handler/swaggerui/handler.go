package swaggerui

import (
	"os"

	"github.com/nurbekabilev/go-open-api/internal/fs"
)

func GetSwaggerYml() []byte {
	yml, err := os.ReadFile(fs.RootPath() + "/api/swagger.yml")
	if err != nil {
		panic("Unable to read swagger yml")
	}

	return yml
}
