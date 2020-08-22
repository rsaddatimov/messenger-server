package api

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

type chatAddRequest struct {
    Name string
    Users []string
}

func(s *ServerAPI) AddChat(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Wrong request method", http.StatusBadRequest)
        return
    }

    var chat chatAddRequest
    if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Проверяем все ли пользователи есть в базе
    count_query := "SELECT count(id) FROM Users WHERE id = ANY($1)"
    var count int = 0
    users_str := "{" + strings.Join(chat.Users, ", ") + "}"

    if err := s.Conn.QueryRow(count_query, users_str).Scan(&count); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count != len(chat.Users) {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Some users from request are not registered!")
        return
    }

    // Проверяем существует ли уже чат с данным названием
    count_query = "SELECT count(id) FROM Chats WHERE name=$1"

    if err := s.Conn.QueryRow(count_query, chat.Name).Scan(&count); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count > 0 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Chat with given name already exists!")
        return
    }

    // Добавляем чам в базу
    const insert_query = "INSERT INTO Chats(name, users) values($1, $2) RETURNING id"
    var id int
    if err := s.Conn.QueryRow(insert_query, chat.Name, users_str).Scan(&id); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    fmt.Fprintln(w, id)
    w.WriteHeader(http.StatusOK)
}