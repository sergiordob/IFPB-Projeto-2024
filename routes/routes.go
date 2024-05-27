package routes

import (
    "github.com/gorilla/mux"
    "github.com/sergiordob/IFPB-Projeto-2024/controllers"
)

func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/drugstores/{estado}", controllers.EndPoint03()).Methods("GET")
    router.HandleFunc("/drugstores/", controllers.EndPoint01()).Methods("GET")
    router.HandleFunc("/teste/", controllers.EndPoint02()).Methods("GET")

    return router
}
