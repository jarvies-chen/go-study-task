// 题目1：使用SQL扩展库进行查询
package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 创建测试数据库和表
func createTestDatabase(db *sqlx.DB) error {
	// 创建表
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			department TEXT NOT NULL,
			salary REAL NOT NULL
			)
	`)
	if err != nil {
		return err
	}

	//插入测试数据
	testData := []Employee{
		{ID: 1, Name: "张三", Department: "技术部", Salary: 15000.00},
		{ID: 2, Name: "李四", Department: "技术部", Salary: 18000.00},
		{ID: 3, Name: "王五", Department: "销售部", Salary: 12000.00},
		{ID: 4, Name: "赵六", Department: "技术部", Salary: 20000.00},
		{ID: 5, Name: "钱七", Department: "人事部", Salary: 10000.00},
		{ID: 6, Name: "孙八", Department: "技术部", Salary: 16000.00},
	}

	//情况数据表
	_, err = db.Exec(`
		DELETE FROM employees
	`)
	if err != nil {
		return err
	}

	//插入数据
	for _, emp := range testData {
		_, err = db.NamedExec(`
			INSERT INTO employees(id, name, department, salary) 
			VALUES (:id, :name, :department, :salary)
		`, emp)
		if err != nil {
			fmt.Println("插入数据失败", err)
		}
	}
	fmt.Println("测试数据初始化完成")
	return nil
}

// 查询技术部所有员工
func queryTechEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee

	//使用 Select 查询多行数据
	query := "SELECT id, name, department, salary FROM employees where department = ?"
	err := db.Select(&employees, query, "技术部")
	if err != nil {
		return nil, fmt.Errorf("查询技术部员工失败: %w", err)
	}
	return employees, nil
}

// 查询技术部所有员工
func queryTechEmployeesNamed(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee

	// 使用 NamedQuery 查询
	query := "SELECT id, name, department, salary FROM employees WHERE department = :dept"
	rows, err := db.NamedQuery(query, map[string]interface{}{"dept": "技术部"})
	if err != nil {
		return nil, fmt.Errorf("NamedQuery 查询失败: %w", err)
	}
	defer rows.Close()

	// 遍历结果
	for rows.Next() {
		var emp Employee
		err := rows.StructScan(emp)
		if err != nil {
			return nil, fmt.Errorf("扫描结果失败: %w", err)
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

// 查询工资最高的员工
func queryHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(employee, query)
	if err != nil {
		return nil, fmt.Errorf("查询最高工资员工失败: %w", err)
	}
	return &employee, nil
}

// 查询工资最高的员工
func queryHighestPaidEmployeeInDepartment(db *sqlx.DB, department string) (*Employee, error) {
	var employee Employee

	// 查询指定部门中工资最高的员工
	query := "SELECT id, name, department, salary FROM employees WHERE department = ? ORDER BY salary DESC LIMIT 1"
	err := db.Get(employee, query, department)
	if err != nil {
		return nil, fmt.Errorf("查询部门最高工资员工失败: %w", err)
	}
	return &employee, nil
}

// 打印员工列表
func printEmployees(employees []Employee, title string) {
	fmt.Printf("\n=== %s ===\n", title)
	if len(employees) == 0 {
		fmt.Println("没有找到员工")
		return
	}
	fmt.Printf("找到 %d 名员工:\n", len(employees))
	fmt.Println("ID\t姓名\t\t部门\t\t工资")
	fmt.Println("----------------------------------------")
	for _, emp := range employees {
		fmt.Printf("%d\t%s\t\t%s\t\t%.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}
}

// 打印单个员工信息
func printEmployee(employee *Employee, title string) {
	fmt.Printf("\n=== %s ===\n", title)
	if employee == nil {
		fmt.Println("没有找到员工")
		return
	}

	fmt.Printf("ID: %d\n", employee.ID)
	fmt.Printf("姓名: %s\n", employee.Name)
	fmt.Printf("部门: %s\n", employee.Department)
	fmt.Printf("工资: %.2f\n", employee.Salary)
}

// 查询所有员工
func queryAllEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees"
	err := db.Select(&employees, query)

	if err != nil {
		return nil, fmt.Errorf("查询所有员工失败: %w", err)
	}
	return employees, nil
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
		log.Fatal("初始化测试数据失败:", err)
	}
	// 查询所有员工
	allEmployees, err := queryAllEmployees(db)
	if err != nil {
		log.Fatal("查询所有员工失败:", err)
	}
	printEmployees(allEmployees, "所有员工")

	// 1. 查询技术部所有员工 - 使用 Select
	fmt.Println("\n--- 方法1：使用 Select 查询技术部员工 ---")
	techEmployees, err := queryTechEmployees(db)
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
	} else {
		printEmployees(techEmployees, "技术部员工")
	}

	// 2. 查询技术部所有员工 - 使用 NamedQuery
	fmt.Println("\n--- 方法2：使用 NamedQuery 查询技术部员工 ---")
	techEmployees2, err := queryTechEmployeesNamed(db)
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
	} else {
		printEmployees(techEmployees2, "技术部员工 (NamedQuery)")
	}

	// 3. 查询工资最高的员工 - 全公司
	fmt.Println("\n--- 查询全公司工资最高的员工 ---")
	highestPaid, err := queryHighestPaidEmployee(db)
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
	} else {
		printEmployee(highestPaid, "全公司工资最高的员工")
	}

	// 4. 查询技术部工资最高的员工
	fmt.Println("\n--- 查询技术部工资最高的员工 ---")
	highestPaidTech, err := queryHighestPaidEmployeeInDepartment(db, "技术部")
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
	} else {
		printEmployee(highestPaidTech, "技术部工资最高的员工")
	}

	// 5. 其他查询示例
	fmt.Println("\n--- 其他查询示例 ---")
	// 查询高薪员工（工资大于15000）
	var highSalaryEmployees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE salary > ? ORDER BY salary DESC"
	err = db.Select(&highSalaryEmployees, query, 15000)
	if err != nil {
		fmt.Printf("查询高薪员工失败: %v\n", err)
	} else {
		printEmployees(highSalaryEmployees, "高薪员工（工资>15000）")
	}
	// 使用 IN 查询多个部门
	var multiDeptEmployees []Employee
	multiDeptQuery := "SELECT id, name, department, salary FROM employees WHERE department IN (?, ?)"
	err = db.Select(&multiDeptEmployees, multiDeptQuery, "技术部", "销售部")
	if err != nil {
		fmt.Printf("多部门查询失败: %v\n", err)
	} else {
		printEmployees(multiDeptEmployees, "技术部和销售部员工")
	}
}
