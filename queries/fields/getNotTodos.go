package queries

import (
	"context"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	database "graphql-mongo/data"
	"graphql-mongo/types"
	"os"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var GetNotTodos = &graphql.Field{
	Type:        graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		databaseName := os.Getenv("DATABASE_NAME")
		collection := os.Getenv("DATABASE_COLLECTION")

		notTodoCollection := database.Client.Database(databaseName).Collection(collection)
		cursor, err := notTodoCollection.Find(context.Background(), bson.D{})

		if err != nil {
			panic(err)
		}

		var todosList []todoStruct

		for cursor.Next(context.Background()) {

			var todo todoStruct
			err := cursor.Decode(&todo)
			if err != nil {
				panic(err)
			}
			todosList = append(todosList, todo)
		}
		return todosList, nil
	},
}
