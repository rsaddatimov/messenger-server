package api

import (
    "database/sql"
    "net/http"
)

type ServerAPI struct {
    Conn *sql.DB
}

func(s *ServerAPI) FetchChats(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        return
    }
}

func(s *ServerAPI) FetchChatsMessages(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        return
    }
}
