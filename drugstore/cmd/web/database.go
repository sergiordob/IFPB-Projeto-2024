package main

import(
	"context"
	"fmt"
	"log"
	"time"
	"github.com/sergiordob/IFPB-Projeto-2024/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	//Conexao com o Banco de Dados

	//String de Configuracao do MondoDBAtlas
	uri := "mongodb+srv://root:root@buscafarma.xprmej7.mongodb.net/?retryWrites=true&w=majority&appName=BuscaFarma"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Configura a opções do cliente
	clientOptions := options.Client().ApplyURI(uri)

	//Se conecta ao cluster
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx) //qualquer coisa, tirar isso

	//Verifica a conexão com o cluster
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB Atlas")
		fmt.Println(err)
	}

	if err == nil {
		fmt.Println("Connected to MongoDB Atlas")
	}

	return client
}

func TestMongoDB(client *mongo.Client) {
	//Acessando o banco de dados 'farma' e listando as collections:

	//Estabelece conexão com o banco de dados 'farma'
	database := client.Database("farma")

	//Printa no terminal os dados da collection 'farma_full_collection'
	collection := database.Collection("farma_full_collection")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Dados
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	//log.Println(len(results))
	for _, result := range results {
		fmt.Println("ID: ", result.ID.Hex())
		fmt.Println("Nome: ", result.Nome)
		fmt.Println("Latitude: ", fmt.Sprintf("%f", result.LongLat[0]))
		fmt.Println("Longitude: ", fmt.Sprintf("%f", result.LongLat[1]))
		fmt.Println("Nome Fantasia: ", result.NomeFantasia)
		fmt.Println("Endereço: ")
		fmt.Println("\tRua: ", result.Endereco.Rua)
		fmt.Println("\tNúmero: ", result.Endereco.Numero)
		fmt.Println("\tBairro: ", result.Endereco.Bairro)
		fmt.Println("\tMunicípio: ", result.Endereco.Municipio)
		fmt.Println("\tEstado: ", result.Endereco.Estado)
		fmt.Println("-----------------------------")
	}
}