package main

import (
	"github.com/sergiordob/IFPB-Projeto-2024/database"
	"github.com/sergiordob/IFPB-Projeto-2024/config"
)



func main() {
	//Testa Banco de Dados
	//TestMongoDB(ConnectMongoDB())

	//Conecta com Banco de Dados
	database.ConnectMongoDB()


	//Inicializa Servidor
	config.InitializeServer()
}
