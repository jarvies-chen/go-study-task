package handler

import (
	"blog/internal/model"
	"blog/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	CommentRepo *repository.CommentRepository
}

func NewCommentHandler(commentRepo *repository.CommentRepository) *CommentHandler {
	return &CommentHandler{CommentRepo: commentRepo}
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	comments, err := h.CommentRepo.FindByPostID(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.UserID = userID

	if err := h.CommentRepo.Create(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}
	c.JSON(http.StatusCreated, comment)
}
