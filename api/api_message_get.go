package api

import (
    "encoding/json"
    "log"
    "net/http"
)

type messageGetRequest struct {
    Chat string
}

type message struct {
    Id         string `json:"id"`
    Chat       string `json:"chat"`
    Author     string `json:"author"`
    Text       string `json:"text"`
    Created_at string `json:"created_at"`
}

func(s *ServerAPI) FetchChatsMessages(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Wrong request method", http.StatusBadRequest)
        return
    }

    var from messageGetRequest
    if err := json.NewDecoder(r.Body).Decode(&from); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Проверяем существует ли данный чат
    if count, err := countEntity(s.Conn, &from.Chat, "Chats"); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count == 0 {
        http.Error(w, "Given chat does not exist!", http.StatusBadRequest)
        return
    }

    // Обращаемся к базе за сообщениями
    const query_str = "SELECT * FROM Messages WHERE chat=$1 ORDER BY created_at"
    rows, err := s.Conn.Query(query_str, from.Chat)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer rows.Close()

    // Конструируем ответ
    msg_array := []message{}
    for rows.Next() {
        cur_msg := message{}

        if err := rows.Scan(
            &cur_msg.Id,
            &cur_msg.Chat,
            &cur_msg.Author,
            &cur_msg.Text,
            &cur_msg.Created_at,
        ); err != nil {
            log.Println(err)
            continue
        }

        msg_array = append(msg_array, cur_msg)
    }

    json.NewEncoder(w).Encode(msg_array)
    w.WriteHeader(http.StatusOK)
}