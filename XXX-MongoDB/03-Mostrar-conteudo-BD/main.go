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

	// Mostrar todo o coneúdo do banco de dados
	cursor, err := client.Database("minhaDatabase").Collection("user").Find(clientContext, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Iterar através do cursor; quando não tiver mais dados o próprio cursor retorna nil para
	// o for e então ele encerra
	for cursor.Next(clientContext) {
		// Variável recipiente do tipo bson.M
		var document bson.M

		// Função decode recebe um endereço de memória para o recipiente da consulta
		err := cursor.Decode(&document)
		if err != nil {
			log.Fatal(err)
		}

		// Printar o resultado
		fmt.Println(document)
	}

	// Fechar a conexão com o servidor
	err = client.Disconnect(clientContext)
	if err != nil {
		log.Fatal(err)
	}
}
