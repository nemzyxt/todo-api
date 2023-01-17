package models

type Task struct {
	Id int `json:"id" gorm:"primary_key auto_increment"`
	Title string `json:"title"`
	Description string `json:"description"`
}