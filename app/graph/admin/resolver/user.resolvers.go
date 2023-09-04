package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/kensei18/enechain-technical-assignment/app/graph/admin"
	"github.com/kensei18/enechain-technical-assignment/app/graph/model"
)

// Company is the resolver for the company field.
func (r *userResolver) Company(ctx context.Context, obj *model.User) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: Company - company"))
}

// User returns admin.UserResolver implementation.
func (r *Resolver) User() admin.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }