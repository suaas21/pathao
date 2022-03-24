package initialize

import (
	"github.com/graphql-go/graphql"
)

func GetUserObject() *graphql.Object {
	// user objects....
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"fullName": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
