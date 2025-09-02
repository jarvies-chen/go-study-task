// 题目1：基本CRUD操作
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint
	Name  string
	Age   int
	Grade string
}

func insert(db *gorm.DB) {
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	db.Create(&student)
}

func search(db *gorm.DB) []Student {
	var students []Student
	db.Where("Age > 18").Find(&students)
	return students
}

func updateGrade(db *gorm.DB) {
	db.Model(&Student{}).Where("Name = ?", "张三").Update("Grade", "四年级")
}

func deleteStudent(db *gorm.DB) {
	db.Where("Age < ?", 15).Delete(&Student{})
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}

	fmt.Println("数据库连接成功")

	db.AutoMigrate(&Student{})

	fmt.Println("数据表创建成功")
	insert(db)

	students := search(db)
	fmt.Println("Found Students ", students)
	updateGrade(db)

	var student Student
	db.Where("name = ?", "张三").First(&student)
	fmt.Println("修改后的年级 ", student.Grade)

	stud := Student{
		Name:  "李四",
		Age:   13,
		Grade: "一年级",
	}
	db.Create(&stud)
	fmt.Println(stud)

	deleteStudent(db)

	allStudents := search(db)
	fmt.Println("Found Students ", allStudents)
	updateGrade(db)
}
