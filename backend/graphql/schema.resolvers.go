package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"imdbv2/ent"
)

func (r *directorResolver) Children(ctx context.Context, obj *ent.Director) ([]*ent.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *movieResolver) Director(ctx context.Context, obj *ent.Movie) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *movieResolver) Parent(ctx context.Context, obj *ent.Movie) (*ent.Director, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovie(ctx context.Context, movie MovieInput) (*ent.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*ent.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

// Director returns DirectorResolver implementation.
func (r *Resolver) Director() DirectorResolver { return &directorResolver{r} }

// Movie returns MovieResolver implementation.
func (r *Resolver) Movie() MovieResolver { return &movieResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type directorResolver struct{ *Resolver }
type movieResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
