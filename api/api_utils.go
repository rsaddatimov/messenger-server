package api

import (
    "fmt"
    "database/sql"
)

func countEntity(conn *sql.DB, id *string, entity_name string) (int, error) {
    query_str := fmt.Sprintf("SELECT count(id) FROM %s WHERE id=$1", entity_name)

    var count int = 0
    err := conn.QueryRow(query_str, id).Scan(&count)

    return count, err
}
