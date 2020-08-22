package api

import (
    "database/sql"
    "fmt"
    "net/http"
)

type ServerAPI struct {
    Conn *sql.DB
}

func(s *ServerAPI) AddUser(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User was added!")
}

func(s *ServerAPI) AddChat(w http.ResponseWriter, r *http.Request) {

}

func(s *ServerAPI) SendMessage(w http.ResponseWriter, r *http.Request) {

}

func(s *ServerAPI) FetchChats(w http.ResponseWriter, r *http.Request) {

}

func(s *ServerAPI) FetchChatsMessages(w http.ResponseWriter, r *http.Request) {

}
