package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"onTime/internal/model"
)

type TeacherRepository struct {
	DB *mongo.Client
}

func (r *TeacherRepository) SignUp(ctx context.Context, teacher model.Teacher) (primitive.ObjectID, error) {
	result, err := r.DB.Database("attendance").Collection("teacher").InsertOne(ctx, teacher)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *TeacherRepository) GetByLogin(ctx context.Context, login string) (model.Teacher, error) {
	var teacher model.Teacher
	filter := bson.M{"login": login}
	err := r.DB.Database("attendance").Collection("teacher").FindOne(ctx, filter).Decode(&teacher)
	if err != nil {
		return model.Teacher{}, err
	}
	return teacher, nil
}

func (r *TeacherRepository) GetClasses(ctx context.Context, teacherLogin string) ([]model.Class, error) {
	filter := bson.M{"teacher_login": teacherLogin}

	cur, err := r.DB.Database("attendance").Collection("class").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var classes []model.Class

	for cur.Next(ctx) {
		var class model.Class
		if err := cur.Decode(&class); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}

func NewTeacherRepository(db *mongo.Client) *TeacherRepository {
	return &TeacherRepository{
		DB: db,
	}
}
