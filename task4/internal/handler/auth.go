package handler

import (
	"blog/internal/model"
	"blog/internal/repository"
	"blog/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo  *repository.UserRepository
	JWTSecret string
	JWTExpire time.Duration
}

func NewAuthHandler(userRepo *repository.UserRepository, secret string, expire time.Duration) *AuthHandler {
	return &AuthHandler{
		UserRepo:  userRepo,
		JWTSecret: secret,
		JWTExpire: expire,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := h.UserRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserRepo.FindByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.Username, h.JWTSecret, h.JWTExpire)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
