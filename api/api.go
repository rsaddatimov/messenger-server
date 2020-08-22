package api

import (
    "database/sql"
)

type ServerAPI struct {
    Conn *sql.DB
}
