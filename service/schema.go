package service

import (
	"fmt"
	"github.com/suaas21/pathao/utils"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/suaas21/pathao/model"
	"github.com/suaas21/pathao/service/initialize"
	"golang.org/x/net/context"
)

func (u *User) UserSchema(incomingReq model.GraphQLIncomingRequest) (*graphql.Result, error) {
	// initialize user object...
	userType := initialize.GetUserObject()

	// query and mutation....
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        userType,
				Description: "get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var user *model.User
					var err error
					if id, ok := p.Args["id"].(string); ok {
						user, err = u.userRepo.GetUser(p.Context, id)
					}

					// set current time and hide password
					if user != nil {
						user.Password = ""
						user.Time = utils.GetCurrentTime()
					}

					return user, err
				},
			},
		},
	})
	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Description: "upsert user",
				Type:        userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userReq := &model.User{}
					var userIDProvided bool
					if val, ok := p.Args["id"].(string); ok {
						userIDProvided = true
						userReq.ID = val
						userReq.Xkey = val
					} else {
						id := uuid.New().String()
						userReq.ID = id
						userReq.Xkey = id
					}

					if val, ok := p.Args["firstName"].(string); ok {
						userReq.FirstName = val
					}

					if val, ok := p.Args["lastName"].(string); ok {
						userReq.LastName = val
					}

					if val, ok := p.Args["password"].(string); ok {
						userReq.Password = val
					}

					// set full name from first and last name
					userReq.FullName = userReq.FirstName + " " + userReq.LastName

					if !userIDProvided {
						err := u.userRepo.CreateUser(p.Context, *userReq)
						if err != nil {
							return nil, err
						}
					} else {
						err := u.userRepo.UpdateUser(p.Context, *userReq)
						if err != nil {
							return nil, err
						}
					}
					user, err := u.userRepo.GetUser(p.Context, userReq.ID)
					if err != nil {
						return nil, err
					}
					if user != nil {
						// return only user id, hide rest of the thing
						user.FirstName = ""
						user.LastName = ""
						user.Password = ""
						user.FullName = ""
					}

					return user, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		fmt.Printf("Failed to create new service, error: %v\n", err)
	}
	ctx := context.WithValue(context.Background(), "loaders", nil)

	return graphql.Do(graphql.Params{
		Context:       ctx,
		Schema:        schema,
		RequestString: incomingReq.Query,
	}), nil

}
