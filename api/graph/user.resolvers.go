package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"

	"github.com/tapestrylabs/users/api/ent"
	"github.com/tapestrylabs/users/api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).
		User.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).
		User.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// RequestLoginCode is the resolver for the requestLoginCode field.
func (r *mutationResolver) RequestLoginCode(ctx context.Context, input model.RequestLoginCode) (string, error) {
	return input.Email, nil
}

// ValidateLoginCode is the resolver for the validateLoginCode field.
func (r *mutationResolver) ValidateLoginCode(ctx context.Context, input model.ValidateLoginCode) (string, error) {
	return input.Code, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
