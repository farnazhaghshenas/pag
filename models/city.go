package models

type City struct {
	Id   int32  `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
}
