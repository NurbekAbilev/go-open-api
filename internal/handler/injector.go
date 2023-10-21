package handler

import "database/sql"

type Injector struct {
	DB *sql.DB
}
