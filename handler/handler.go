package handler

import (
    "context"
    "encoding/json"
    "net/http"

    "IFPB-Projeto-2024/model"
    "IFPB-Projeto-2024/mongo"

    "go.mongodb.org/mongo-driver/bson"
)

func GetDados(w http.ResponseWriter, r *http.Request) {
    client := mongo.Connect()
    collection := client.Database("farma").Collection("farma_full_collection")

    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var results []model.Dados
    if err = cursor.All(context.TODO(), &results); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}
