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

func (cr *categoryRepository) FetchById(c context.Context, id string) (*domain.Category, error) {
	collection := cr.database.Collection(cr.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var category *domain.Category

	err = collection.FindOne(c, bson.M{
		"_id": idHex,
	}).Decode(&category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *categoryRepository) Create(c context.Context, category *domain.Category) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, category)
	return err
}

func (cr *categoryRepository) Fetch(c context.Context, query domain.CategorySearch) ([]domain.Category, error) {
	collection := cr.database.Collection(cr.collection)
	var filter bson.A
	filter = bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"$and",
						bson.A{
							bson.D{{"type", query.Type}},
							bson.D{{"level", query.Level}},
						},
					},
				},
			},
		},
	}
	if query.Level == 0 {
		filter = bson.A{
			bson.D{
				{"$match",
					bson.D{{"type", query.Type}},
				},
			},
		}
	}

	cursor, err := collection.Aggregate(c, filter)
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
