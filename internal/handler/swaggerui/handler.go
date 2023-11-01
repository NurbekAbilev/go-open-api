package swaggerui

import "os"

func GetSwaggerYml() []byte {
	yml, err := os.ReadFile("../../api/swagger.yml")
	if err != nil {
		panic("Unable to read swagger yml")
	}

	return yml
}
