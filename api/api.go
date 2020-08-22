package api

import (
    "database/sql"
    "net/http"
)

type ServerAPI struct {
    Conn *sql.DB
}

func(s *ServerAPI) AddChat(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        wrongMethod(w, r)
        return
    }
}

func(s *ServerAPI) SendMessage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        wrongMethod(w, r)
        return
    }
}

func(s *ServerAPI) FetchChats(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        wrongMethod(w, r)
        return
    }
}

func(s *ServerAPI) FetchChatsMessages(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        wrongMethod(w, r)
        return
    }
}
