package queries

import (
	"github.com/graphql-go/graphql"
	queries "graphql-mongo/queries/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getNotTodos": queries.GetNotTodos,
	},
})
