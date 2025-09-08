package repository

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.DB.Create(comment).Error
}

func (r *CommentRepository) FindByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error
	return comments, err
}
