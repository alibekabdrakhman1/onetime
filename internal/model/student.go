package model

type Student struct {
	Login    string  `bson:"login"`
	Password string  `bson:"password"`
	Name     string  `bson:"name"`
	Classes  []Class `bson:"classes"`
}
type LogIn struct {
	Login    string
	Password string
}
type StudentCreate struct {
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
}
