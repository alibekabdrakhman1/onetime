package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"onTime/internal/model"
)

type StudentRepository struct {
	DB *mongo.Client
}

func (r *StudentRepository) GetClasses(ctx context.Context, studentLogin string) ([]model.Class, error) {
	var student model.Student
	filter := bson.M{"login": studentLogin}
	err := r.DB.Database("attendance").Collection("class").FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student.Classes, nil
}

func (r *StudentRepository) SignUp(ctx context.Context, student model.Student) (primitive.ObjectID, error) {
	result, err := r.DB.Database("attendance").Collection("student").InsertOne(ctx, student)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *StudentRepository) GetByLogin(ctx context.Context, login string) (model.Student, error) {
	var student model.Student
	filter := bson.M{"login": login}
	err := r.DB.Database("attendance").Collection("class").FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return model.Student{}, err
	}
	return student, nil
}

func (r *StudentRepository) Attend(ctx context.Context, classId string, studentId string) error {
	filterC := bson.M{"_id": classId}
	updateC := bson.M{"$set": bson.M{fmt.Sprintf("map.%s", studentId): true}}
	result, err := r.DB.Database("attendance").Collection("class").UpdateOne(ctx, filterC, updateC)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no document found to update")
	}
	return nil
}

func NewStudentRepository(db *mongo.Client) *StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}
