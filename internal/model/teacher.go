package model

type Teacher struct {
	Login    string   `bson:"login"`
	Password string   `bson:"password"`
	Name     string   `bson:"name"`
	Classes  []string `bson:"classes"`
}
type TeacherCreate struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
}
