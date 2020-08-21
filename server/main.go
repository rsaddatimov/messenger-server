package main

import (
	"database/sql"
	"log"
	"net/http"

	"./api"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres")

	if err != nil {
		log.Fatal(err)
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
