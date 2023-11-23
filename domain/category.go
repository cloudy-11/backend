package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCategory = "categories"
)

type CategorySearch struct {
	Type  string `form:"type" binding:"omitempty"`
	Level int8   `form:"level" binding:"omitempty"`
	//Name  string `form:"name" binding:"omitempty"`
}

type Category struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Type   string             `bson:"type" form:"type" biding:"required" json:"type"`    // Coding | English
	Slug   string             `bson:"slug" json:"slug"`                                  // Coding | English
	Level  int8               `bson:"level" form:"level" biding:"required" json:"level"` // Elementary | Intermediate | Advance
	Name   string             `bson:"name" form:"name" biding:"required" json:"name"`
	IsLock bool               `bson:"is_lock" form:"isLock,default=true" json:"isLock"`
}

type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	FetchById(c context.Context, id string) (*Category, error)
	Fetch(c context.Context, query CategorySearch) ([]Category, error)
	Delete(c context.Context, id string) error
}

type CategoryUseCase interface {
	Create(c context.Context, category *Category) error
	FetchById(c context.Context, id string) (*Category, error)
	Fetch(c context.Context, query CategorySearch) ([]Category, error)
	Delete(c context.Context, id string) error
}
