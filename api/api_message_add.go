package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type messageAddRequest struct {
    Chat string
    Author string
    Text string
}

func(s *ServerAPI) SendMessage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Wrong request method", http.StatusBadRequest)
        return
    }

    var message messageAddRequest
    if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Проверяем существует ли данный чат
    if count, err := countEntity(s.Conn, &message.Chat, "Chats"); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count == 0 {
        http.Error(w, "Given chat does not exist!", http.StatusBadRequest)
        return
    }

    // Проверяем существует ли данный пользователь
    if count, err := countEntity(s.Conn, &message.Author, "Users"); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count == 0 {
        http.Error(w, "Given user is not registered!", http.StatusBadRequest)
        return
    }

    // Состоит ли данный пользователь в данном чате
    query := "SELECT count(id) FROM Chats WHERE id=$1 AND $2=ANY(users)"
    var count int
    if err := s.Conn.QueryRow(query, message.Chat, message.Author).Scan(&count); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count == 0 {
        http.Error(w, "Given user is not a member of given chat!", http.StatusForbidden)
        return
    }

    // Записываем сообщение
    query = "INSERT INTO Messages(chat, author, text) values($1, $2, $3) RETURNING id"
    var id int
    err := s.Conn.QueryRow(query, message.Chat, message.Author, message.Text).Scan(&id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, id)
    w.WriteHeader(http.StatusOK)
}