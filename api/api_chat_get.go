package api

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/lib/pq"
)

type chatGetRequest struct {
    User string
}

type chat struct {
    Id          string `json:"id"`
    Name        string `json:"name"`
    Users     []string `json:"users"`
    Created_at  string `json:"created_at"`
    LastUpdated string `json:"lastUpdated"`
}

func(s *ServerAPI) FetchChats(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Wrong request method", http.StatusBadRequest)
        return
    }

    var user chatGetRequest
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Проверяем существование данного пользователя
    if count, err := countEntity(s.Conn, &user.User, "Users"); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    } else if count == 0 {
        http.Error(w, "Given user does not exist!", http.StatusBadRequest)
        return
    }

    // Обращаемся к базе за чатами
    const query_str = "SELECT * from Chats WHERE $1=ANY(users) ORDER BY lastUpdated DESC"
    rows, err := s.Conn.Query(query_str, user.User)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer rows.Close()

    // Конструируем ответ
    chats_array := []chat{}
    for rows.Next() {
        cur_chat := chat{}

        if err := rows.Scan(
            &cur_chat.Id,
            &cur_chat.Name,
            pq.Array(&cur_chat.Users),
            &cur_chat.Created_at,
            &cur_chat.LastUpdated,
        ); err != nil {
            log.Println(err)
            continue
        }
        
        chats_array = append(chats_array, cur_chat)
    }

    json.NewEncoder(w).Encode(chats_array)
    w.WriteHeader(http.StatusOK)
}
