package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type userRequest struct {
    Username string
}

func(s *ServerAPI) AddUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Wrong request method", http.StatusBadRequest)
        return
    }

    var user userRequest
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Проверяем существует ли уже данный пользователь в базе
    const count_query = "SELECT count(id) FROM Users WHERE username=$1"
    var count int = 0
    
    if err := s.Conn.QueryRow(count_query, user.Username).Scan(&count); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count > 0 {
        errstr := fmt.Sprintf("%s is already registered!", user.Username)
        http.Error(w, errstr, http.StatusNotFound)
        return
    }
    
    // Записываем пользователя
    const insert_query = "INSERT INTO Users(username) values($1) RETURNING id"
    var id int
    if err := s.Conn.QueryRow(insert_query, user.Username).Scan(&id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, id)
    w.WriteHeader(http.StatusOK)
}
