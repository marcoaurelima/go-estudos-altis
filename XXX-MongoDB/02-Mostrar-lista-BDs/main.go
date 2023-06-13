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
	// Configurar o context
	clientContext := context.Background()

	// Configurar as opções do cliente
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectar ao servidor MongoDB
	client, err := mongo.Connect(clientContext, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar a conexão com o servidor
	err = client.Ping(clientContext, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Listar os nomes dos bancos de dados
	databaseNames, err := client.ListDatabaseNames(clientContext, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Imprimir os nomes dos bancos de dados
	for _, name := range databaseNames {
		fmt.Println(name)
	}

	// Fechar a conexão com o servidor
	err = client.Disconnect(clientContext)
	if err != nil {
		log.Fatal(err)
	}
}
