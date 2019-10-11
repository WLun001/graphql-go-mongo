package mutations

import (
	"github.com/graphql-go/graphql"
	mutations "graphql-mongo/mutations/fields"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": mutations.CreateNotTodo,
	},
})
