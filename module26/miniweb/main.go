package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", GetTime)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from: %s\n", r.RemoteAddr)
	_, err := fmt.Fprintf(w, "<h1>%s</h1><br>", "Запуск приложения в контейнере Docker")
	if err != nil {
		log.Printf("Error: %s\n", err)
	}

	_, err = fmt.Fprintf(w, "Текущие дата и время %s\n",
		time.Now().Format("02-01-2006 15:04"))
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}
