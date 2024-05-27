package main

import (

	//"github.com/sergiordob/IFPB-Projeto-2024/internal/database"
	"github.com/sergiordob/IFPB-Projeto-2024/database"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"github.com/sergiordob/IFPB-Projeto-2024/config"
)

type application struct {
	mongoClient *mongo.Client
	context     context.Context
}

func main() {
	//Testa Banco de Dados
	//TestMongoDB(ConnectMongoDB())

	//Conecta com Banco de Dados
	database.ConnectMongoDB()


	//Inicializa Servidor
	config.InitializeServer()
}
