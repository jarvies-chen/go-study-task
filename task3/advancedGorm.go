// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	PostCount int       `gorm:"default:0"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	User          User      `gorm:"foreignKey:UserID"`
	CommentCount  int       `gorm:"default:0"`
	CommentStatus string    `gorm:"default:'有评论'"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	PostID  uint

	User User `gorm:"foreignKey:UserID"`
	Post Post `gorm:"foreignKey:PostID"`
}

func createTestData(db *gorm.DB) error {
	// 创建用户
	users := []User{
		{Username: "alice", Password: "password123", Email: "alice@example.com"},
		{Username: "bob", Password: "password456", Email: "bob@example.com"},
		{Username: "charlie", Password: "password789", Email: "charlie@example.com"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("创建用户失败: %w", err)
		}
		fmt.Printf("创建用户: %s (ID: %d)\n", user.Username, user.ID)
	}

	// 创建文章
	posts := []Post{
		{Title: "Go语言入门", Content: "Go语言是一门现代化的编程语言...", UserID: 1},
		{Title: "GORM使用指南", Content: "GORM是Go语言的ORM框架...", UserID: 1},
		{Title: "Web开发实践", Content: "使用Go语言开发Web应用...", UserID: 2},
		{Title: "数据库设计", Content: "数据库设计的最佳实践...", UserID: 3},
	}

	for _, post := range posts {
		if err := db.Create(&post).Error; err != nil {
			return fmt.Errorf("创建文章失败: %w", err)
		}
		fmt.Printf("创建文章: %s (ID: %d)\n", post.Title, post.ID)
	}

	// 创建评论
	comments := []Comment{
		{Content: "很好的文章，学到了很多！", UserID: 2, PostID: 1},
		{Content: "感谢分享，期待更多内容", UserID: 3, PostID: 1},
		{Content: "这个例子很清晰易懂", UserID: 2, PostID: 2},
		{Content: "有没有更详细的教程？", UserID: 3, PostID: 2},
		{Content: "实用的开发经验", UserID: 1, PostID: 3},
	}

	for _, comment := range comments {
		if err := db.Create(&comment).Error; err != nil {
			return fmt.Errorf("创建评论失败: %w", err)
		}
		fmt.Printf("创建评论: %s (ID: %d)\n", comment.Content[:20]+"...", comment.ID)
	}

	return nil
}

func queryExamples(db *gorm.DB) {
	fmt.Println("1. 查询所有用户及其文章数量")
	type UserPostCount struct {
		Username  string
		PostCount int
	}
	var userPostCounts []UserPostCount
	db.Model(&User{}).Select("users.username, count(posts.id) as post_count").
		Joins("left join posts on posts.user_id = users.id").
		Group("users.id").
		Scan(&userPostCounts)
	for _, upc := range userPostCounts {
		fmt.Printf("  用户: %s, 文章数: %d\n", upc.Username, upc.PostCount)
	}

	fmt.Println("2. 查询特定用户的所有文章")
	var user User
	if err := db.Preload("Posts").Where("username = ?", "alice").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("  用户不存在")
		} else {
			fmt.Printf("  查询失败: %v\n", err)
		}
		return
	}

	fmt.Printf("  用户: %s (ID: %d)\n", user.Username, user.ID)
	fmt.Printf("  文章数量: %d\n", len(user.Posts))
	for _, post := range user.Posts {
		fmt.Printf("    - %s\n", post.Title)
	}

	fmt.Println("3. 查询文章及其作者和评论")
	var post Post
	if err := db.Preload("User").Preload("Comments.User").Where("title = ?", "Go语言入门").First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("  文章不存在")
		} else {
			fmt.Printf("  查询失败: %v\n", err)
		}
		return
	}
	fmt.Printf("  标题: %s\n", post.Title)
	fmt.Printf("  作者: %s\n", post.User.Username)
	fmt.Printf("  评论数量: %d\n", len(post.Comments))
	for _, comment := range post.Comments {
		fmt.Printf("    - %s (评论者: %s)\n", comment.Content, comment.User.Username)
	}

	// 4. 查询所有文章及其作者
	fmt.Println("\n4. 所有文章及其作者:")
	var posts []Post
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		fmt.Printf("  查询失败: %v\n", err)
		return
	}

	for _, post := range posts {
		fmt.Printf("  - %s (作者: %s)\n", post.Title, post.User.Username)
	}

	// 5. 查询用户的评论
	fmt.Println("\n5. 用户 Bob 的所有评论:")
	var user2 User
	if err := db.Preload("Comments.Post").Where("username = ?", "bob").First(&user2).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("  用户不存在")
		} else {
			fmt.Printf("  查询失败: %v\n", err)
		}
		return
	}

	fmt.Printf("  用户: %s\n", user2.Username)
	fmt.Printf("  评论数量: %d\n", len(user2.Comments))
	for _, comment := range user2.Comments {
		fmt.Printf("    - %s (文章: %s)\n", comment.Content, comment.Post.Title)
	}
}

// 钩子函数实现

// BeforeCreate 在创建文章前的钩子函数
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	fmt.Printf("文章创建前钩子: 文章 '%s' 即将创建\n", p.Title)
	return nil
}

// AfterCreate 在创建文章后的钩子函数 - 更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Printf("文章创建后钩子: 更新用户 %d 的文章数量\n", p.UserID)

	// 更新用户的文章数量
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error; err != nil {
		return fmt.Errorf("更新用户文章数量失败: %w", err)
	}

	fmt.Printf("用户 %d 的文章数量已更新\n", p.UserID)
	return nil
}

// BeforeDelete 在删除评论前的钩子函数
func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("评论删除前钩子: 评论 %d 即将删除\n", c.ID)
	return nil
}

// AfterDelete 在删除评论后的钩子函数 - 检查并更新文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Printf("评论删除后钩子: 检查文章 %d 的评论状态\n", c.PostID)

	// 查询该文章剩余的评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return fmt.Errorf("查询文章评论数量失败: %w", err)
	}

	fmt.Printf("文章 %d 剩余评论数量: %d\n", c.PostID, commentCount)

	// 根据评论数量更新文章状态
	var updateData map[string]interface{}
	if commentCount == 0 {
		updateData = map[string]interface{}{
			"comment_count":  commentCount,
			"comment_status": "无评论",
		}
		fmt.Printf("文章 %d 评论数量为0，更新状态为 '无评论'\n", c.PostID)
	} else {
		updateData = map[string]interface{}{
			"comment_count":  commentCount,
			"comment_status": "有评论",
		}
		fmt.Printf("文章 %d 更新评论数量为 %d\n", c.PostID, commentCount)
	}

	// 更新文章的评论数量和状态
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(updateData).Error; err != nil {
		return fmt.Errorf("更新文章评论状态失败: %w", err)
	}

	return nil
}

// BeforeUpdate 在更新文章前的钩子函数
func (p *Post) BeforeUpdate(tx *gorm.DB) error {
	fmt.Printf("文章更新前钩子: 文章 %d 即将更新\n", p.ID)
	return nil
}

// AfterUpdate 在更新文章后的钩子函数
func (p *Post) AfterUpdate(tx *gorm.DB) error {
	fmt.Printf("文章更新后钩子: 文章 %d 已更新\n", p.ID)
	return nil
}

func testCommentHooks(db *gorm.DB) {
	fmt.Println("\n=== 测试评论钩子函数 ===")

	// 1. 先创建一个只有1条评论的文章
	fmt.Println("1. 创建测试文章和评论:")
	testPost := Post{
		Title:   "测试评论钩子",
		Content: "用于测试评论删除钩子的文章...",
		UserID:  1,
	}

	if err := db.Create(&testPost).Error; err != nil {
		fmt.Printf("创建测试文章失败: %v\n", err)
		return
	}

	testComment := Comment{
		Content: "这是唯一的评论",
		UserID:  2,
		PostID:  testPost.ID,
	}

	if err := db.Create(&testComment).Error; err != nil {
		fmt.Printf("创建测试评论失败: %v\n", err)
		return
	}

	// 更新文章的初始状态
	db.Model(&testPost).Updates(map[string]interface{}{
		"comment_count":  1,
		"comment_status": "有评论",
	})

	fmt.Printf("创建测试文章 (ID: %d) 和评论 (ID: %d)\n", testPost.ID, testComment.ID)

	// 显示当前状态
	var post Post
	db.First(&post, testPost.ID)
	fmt.Printf("删除前文章状态 - 评论数: %d, 状态: %s\n", post.CommentCount, post.CommentStatus)

	// 2. 删除唯一的评论，触发钩子函数
	fmt.Println("\n2. 删除唯一评论，测试钩子函数:")
	if err := db.Delete(&testComment).Error; err != nil {
		fmt.Printf("删除评论失败: %v\n", err)
		return
	}

	fmt.Printf("评论 %d 已删除\n", testComment.ID)

	// 检查文章状态是否更新
	db.First(&post, testPost.ID)
	fmt.Printf("删除后文章状态 - 评论数: %d, 状态: %s\n", post.CommentCount, post.CommentStatus)

	// 3. 测试删除文章的最后一个评论
	fmt.Println("\n3. 测试删除文章的多个评论:")

	// 创建文章和多个评论
	multiPost := Post{
		Title:   "多评论测试",
		Content: "用于测试多个评论删除的文章...",
		UserID:  2,
	}
	db.Create(&multiPost)

	var createdComments []Comment
	comments := []Comment{
		{Content: "评论1", UserID: 1, PostID: multiPost.ID},
		{Content: "评论2", UserID: 3, PostID: multiPost.ID},
		{Content: "评论3", UserID: 1, PostID: multiPost.ID},
	}

	for _, comment := range comments {
		db.Create(&comment)
		createdComments = append(createdComments, comment)
	}

	// 更新文章状态
	db.Model(&multiPost).Updates(map[string]interface{}{
		"comment_count":  3,
		"comment_status": "有评论",
	})

	fmt.Printf("创建多评论文章 (ID: %d) 和 3 条评论\n", multiPost.ID)

	// 删除一条评论
	var firstComment Comment
	db.First(&firstComment, createdComments[0].ID)
	if err := db.Delete(&firstComment).Error; err != nil {
		fmt.Printf("删除评论失败: %v\n", err)
		return
	}

	// 检查文章状态
	var updatedPost Post
	db.First(&updatedPost, multiPost.ID)
	fmt.Printf("删除一条评论后 - 评论数: %d, 状态: %s\n", updatedPost.CommentCount, updatedPost.CommentStatus)

	// 删除剩余评论
	var remainingComments []Comment
	db.Where("post_id = ?", multiPost.ID).Find(&remainingComments)

	for _, comment := range remainingComments {
		db.Delete(&comment)
		fmt.Printf("删除评论 %d\n", comment.ID)
	}

	// 最终检查文章状态
	db.First(&updatedPost, multiPost.ID)
	fmt.Printf("删除所有评论后 - 评论数: %d, 状态: %s\n", updatedPost.CommentCount, updatedPost.CommentStatus)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err := createTestData(db); err != nil {
		fmt.Println("初始化数据失败")
	}
	queryExamples(db)
	testCommentHooks(db)
}
