package mutations

import (
	"context"
	"github.com/graphql-go/graphql"
	database "graphql-mongo/data"
	"graphql-mongo/types"
	"os"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var CreateNotTodo = &graphql.Field{
	Type:        types.NotTodo,
	Description: "Create a not Todo",

	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)

		databaseName := os.Getenv("DATABASE_NAME")
		collection := os.Getenv("DATABASE_COLLECTION")
		notTodoCollection := database.Client.Database(databaseName).Collection(collection)
		_, err := notTodoCollection.InsertOne(context.Background(), map[string]string{"name": name, "description": description})
		if err != nil {
			panic(err)
		}

		return todoStruct{name, description}, nil
	},
}
