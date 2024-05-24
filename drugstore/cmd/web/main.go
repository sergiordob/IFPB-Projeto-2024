package main

import(
	"net/http"
	"log"
)


func main() {

	TestMongoDB(ConnectMongoDB())

	server := http.Server{
		Addr: ":8080",
		Handler: Routes(),
	}

	log.Println("Starting server at localhost:8080...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
