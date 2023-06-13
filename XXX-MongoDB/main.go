package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configurar as opções do cliente
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectar ao servidor MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("1", err)
	}

	// Verificar a conexão com o servidor
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("2", err)
	}

	// Listar os nomes dos bancos de dados
	databaseNames, err := client.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		log.Fatal("<<3>> ", err)
	}

	// Imprimir os nomes dos bancos de dados
	for _, name := range databaseNames {
		fmt.Println(name)
	}

	// Fechar a conexão com o servidor
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("4", err)
	}
}
