//题目2：实现类型安全映射

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

// 创建测试数据库和表
func createTestDatabase(db *sqlx.DB) error {
	db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			price REAL NOT NULL
			)
	`)

	//清空数据库
	db.Exec(`
		DELETE FROM books
	`)

	//创建测试数据
	books := []Book{
		{ID: 1, Title: "水浒传", Author: "施耐庵", Price: 100.00},
		{ID: 2, Title: "红楼梦", Author: "曹雪芹", Price: 50.00},
	}

	for _, book := range books {
		db.NamedExec(`
			INSERT INTO books(id, title, author, price) 
			VALUES (:id, :title, :author, :price)
		`, book)
	}
	return nil
}

func main() {
	//连接数据库
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()

	fmt.Println("数据库连接成功")

	if err := createTestDatabase(db); err != nil {
		fmt.Println("初始化失败")
	}

	//询价格大于 50 元的书籍
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err = db.Select(&books, query, 50)
	if err != nil {
		fmt.Println("查询失败")
	}
	fmt.Println("查询成功")
	fmt.Println(books)
}
