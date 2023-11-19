package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCategory = "categories"
)

type Category struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Type  string             `bson:"type" form:"type" biding:"required"`   // Coding | English
	Level int8               `bson:"level" form:"level" biding:"required"` // Elementary | Intermediate | Advance
	Name  string             `bson:"name" form:"name" biding:"required"`
}

type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	FetchById(c context.Context, id string) (*Category, error)
	Fetch(c context.Context) ([]Category, error)
	Delete(c context.Context, id string) error
}

type CategoryUseCase interface {
	Create(c context.Context, category *Category) error
	FetchById(c context.Context, id string) (*Category, error)
	Fetch(c context.Context) ([]Category, error)
	Delete(c context.Context, id string) error
}
