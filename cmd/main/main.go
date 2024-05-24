package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "IFPB-Projeto-2024/handler"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/dados", handler.GetDados).Methods("GET")

    http.Handle("/", r)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
