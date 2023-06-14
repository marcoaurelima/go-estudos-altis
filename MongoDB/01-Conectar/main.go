package main

import (
	"context"
	"log"

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

	// Fechar a conexão com o servidor
	err = client.Disconnect(clientContext)
	if err != nil {
		log.Fatal(err)
	}
}
