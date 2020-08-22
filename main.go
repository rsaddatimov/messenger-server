package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "time"

    "./api"

    _ "github.com/lib/pq"
)

func Connect() *sql.DB {
    credential := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        "database",
        "5432",
        "postgres",
        "postgres",
        "messengerdb",
        "disable",
    )
    const retries = 5

    for i := 1; i <= retries; i++ {
        log.Printf("Trying to establish connection with database. Try count = %d...\n", i)

        db, err := sql.Open("postgres", credential)

        if err != nil {
            log.Println(err)
            time.Sleep(5 * time.Second)
            continue
        }
        
        if err := db.Ping(); err != nil {
            log.Println(err)
            time.Sleep(5 * time.Second)
        } else {
            return db;
        }
    }

    return nil
}

func main() {
    db := Connect()

    if db == nil {
        log.Fatal("Unable to establish connection to database")
    }
    defer db.Close()

    server := api.ServerAPI{Conn: db}

    http.HandleFunc("/chats/add", server.AddChat)
    http.HandleFunc("/chats/get", server.FetchChats)
    http.HandleFunc("/messages/add", server.SendMessage)
    http.HandleFunc("/messages/get", server.FetchChatsMessages)
    http.HandleFunc("/users/add", server.AddUser)

    http.ListenAndServe(":9000", nil)
}
