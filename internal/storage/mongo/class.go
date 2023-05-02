package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"onTime/internal/model"
)

type ClassRepository struct {
	DB *mongo.Client
}

func (r *ClassRepository) GetAllClasses(ctx context.Context) ([]model.Class, error) {
	cursor, err := r.DB.Database("attendance").Collection("class").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var classes []model.Class
	for cursor.Next(ctx) {
		var class model.Class
		err := cursor.Decode(&class)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepository) Create(ctx context.Context, class model.Class, students []string) (string, error) {
	res, err := r.DB.Database("attendance").Collection("class").InsertOne(ctx, class)
	for i := 0; i < len(students); i++ {
		filter := bson.M{"login": students[i]}
		fmt.Println(students[i])
		update := bson.M{
			"$push": bson.M{
				"classes": class,
			},
		}

		r, err := r.DB.Database("attendance").Collection("student").UpdateOne(ctx, filter, update)
		if r.MatchedCount == 0 {
			fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAA")
		}

		if err != nil {
			return "", err
		}
	}

	return res.InsertedID.(string), err
}

func (r *ClassRepository) GetById(ctx context.Context, id string) (model.Class, error) {
	filter := bson.M{"_id": id}
	var class model.Class
	err := r.DB.Database("attendance").Collection("class").FindOne(ctx, filter).Decode(&class)
	if err != nil {
		return model.Class{}, err
	}
	return class, nil
}

func NewClassRepository(db *mongo.Client) *ClassRepository {
	return &ClassRepository{
		DB: db,
	}
}
