package post

import (
	"echoV1/model"
	"log"

	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) []model.Post {
	var post []model.Post
	db.Order("id asc").Find(&post)
	return post
}

func CreatePost(db *gorm.DB, post *model.Post) *model.Post {
	db.Create(&post)
	return post
}

func GetById[T any](db *gorm.DB, id T) *model.Post {
	var post model.Post
	db.First(&post, id)
	return &post
}

func DeleteById[T any](db *gorm.DB, id T) {
	result := db.Delete(&model.Post{}, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func UpdateById[T any](db *gorm.DB, id T, post model.Post) *model.Post {
	var postUpdate model.Post
	postmodify := db.First(&postUpdate, id)
	if postmodify.Error != nil {
		log.Fatal(postmodify.Error)
	}
	postUpdate = post
	postSave := db.Save(&postUpdate)

	if postSave.Error != nil {
		log.Fatal(postSave.Error)
	}

	return &postUpdate
}
