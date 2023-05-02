package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Class struct {
	Id           primitive.ObjectID `bson:"_id"`
	TeacherLogin string             `bson:"teacher_login"`
	Students     map[string]bool    `bson:"students"`
	StartTime    time.Time          `bson:"start_time"`
	EndTime      time.Time          `bson:"end_time"`
	SecretKey    string             `bson:"secret_key"`
}
type CreateClass struct {
	Students  []string  `json:"students" bson:"students"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
	SecretKey string    `json:"secret_key"`
}
