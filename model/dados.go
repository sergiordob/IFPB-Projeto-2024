package model

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Dados struct {
    ID           primitive.ObjectID `bson:"_id" json:"id"`
    Nome         string             `bson:"nome" json:"nome"`
    LongLat      []float64          `bson:"longLat" json:"longLat"`
    NomeFantasia string             `bson:"nomeFantasia" json:"nomeFantasia"`
    Endereco     struct {
        Rua       string `bson:"rua" json:"rua"`
        Numero    string `bson:"numero" json:"numero"`
        Bairro    string `bson:"bairro" json:"bairro"`
        Municipio string `bson:"municipio" json:"municipio"`
        Estado    string `bson:"estado" json:"estado"`
    } `bson:"endereco" json:"endereco"`
}
