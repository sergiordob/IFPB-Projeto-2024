package mongo

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
    uri := "mongodb+srv://root:root@buscafarma.xprmej7.mongodb.net/?retryWrites=true&w=majority&appName=BuscaFarma"

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB Atlas:", err)
    }

    return client
}
