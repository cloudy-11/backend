package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionQuestion = "questions"
)

type QuestionSearch struct {
	CategoryId string `form:"categoryId" binding:"omitempty"`
	//Name  string `form:"name" binding:"omitempty"`
}
type Question struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	CategoryId  string             `bson:"category_id" form:"categoryId" biding:"required" json:"categoryId"`
	Title       string             `bson:"title" form:"title" biding:"required" json:"title"`
	Description string             `bson:"description" form:"description" biding:"required" json:"description"`
	Examples    []string           `bson:"examples" form:"examples" biding:"required" json:"examples"`
	Level       int8               `bson:"level" form:"level" biding:"required" json:"level"` // Elementary | Intermediate | Advance
	IsLock      bool               `bson:"is_lock" form:"isLock,default=True" json:"isLock"`
	Slug        string             `bson:"slug" json:"slug"`
	Ghost       int32              `bson:"ghost" form:"ghost" biding:"required" json:"ghost"`
}

type QuestionRepository interface {
	Create(c context.Context, question *Question) error
	FetchById(c context.Context, id string) (*Question, error)
	Fetch(c context.Context, query QuestionSearch) ([]Question, error)
	Delete(c context.Context, id string) error
}

type QuestionUseCase interface {
	Create(c context.Context, question *Question) error
	FetchById(c context.Context, id string) (*Question, error)
	Fetch(c context.Context, query QuestionSearch) ([]Question, error)
	Delete(c context.Context, id string) error
}
