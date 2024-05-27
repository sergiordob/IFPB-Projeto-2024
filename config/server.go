package config

import (
    "log"
    "net/http"
    "github.com/sergiordob/IFPB-Projeto-2024/routes"
)

func InitializeServer() {
    router := routes.InitializeRoutes()
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    log.Println("Starting server at localhost:8080...")
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
