package repository

import (
	"context"
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type categoryRepository struct {
	database   mongo.Database
	collection string
}

func (cr *categoryRepository) Create(c context.Context, category *domain.Category) error {
	collection := cr.database.Collection(cr.collection)

	_, err := collection.InsertOne(c, category)

	return err
}

func (cr *categoryRepository) Fetch(c context.Context) ([]domain.Category, error) {
	collection := cr.database.Collection(cr.collection)
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}

	var categories []domain.Category

	err = cursor.All(c, &categories)

	if categories == nil {
		return []domain.Category{}, err
	}

	return categories, nil
}

func (cr *categoryRepository) Delete(c context.Context, id string) error {
	collection := cr.database.Collection(cr.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{
		"_id": idHex,
	})

	return err
}

func NewCategoryRepository(db mongo.Database, collection string) domain.CategoryRepository {
	return &categoryRepository{
		database:   db,
		collection: collection,
	}
}
