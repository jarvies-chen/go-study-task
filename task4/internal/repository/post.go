package repository

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Preload("User").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("User").First(&post, id).Error
	return &post, err
}

func (r *PostRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Post{}, id).Error
}

func (r *PostRepository) FindByUserID(userID uint) ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}
