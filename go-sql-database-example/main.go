// file: main.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/go-sql-driver/mysql"
)

// 第二步:测试curd
var db *sql.DB //定义全局对象
type user struct {
	id   int
	age  int
	name string
}

func main() {
	//pingDemo()
	if err := initDB(); err != nil {
		panic(err)
	}
	//queryRowDemo(3)
	//queryMultiRowDemo()
	//insertRowDemo()
	//updateRowDemo()
	//deleteRowDemo()
	//timeOutDemo()
	//transactionDemo()
	prepareQueryDemo()
}

// 2.0 定义初始化数据库函数
func initDB() (err error) {
	dsn := "coopers:2019Youth@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn) //注意使用全局对象进行赋值
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)
	fmt.Println("init succ")
	return nil
}

// 2.1 查询单条数据
func queryRowDemo(id int) {
	sqlStr := "select id,name,age from user where id = ? limit 1"
	// `QueryRow` always returns a single row from the database
	row := db.QueryRow(sqlStr, id)

	var u user
	if err := row.Scan(&u.id, &u.name, &u.age); err != nil {
		log.Fatalf("could not scan row: %v", err)
	}
	fmt.Printf("user:%+v\n", u)
}

// 2.2 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "select id,name,age from user where id > ?"
	// `QueryRow` always returns a single row from the database
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
	//循环读取结果集数据
	users := []user{}
	for rows.Next() {
		var u user
		// create an instance of `user` and write the result of the current row into it
		if err := rows.Scan(&u.id, &u.name, &u.age); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		users = append(users, u)
	}
	// print the length, and all the birds
	fmt.Printf("len:%v,users:%+v\n", len(users), users)
}

// 插入数据
func insertRowDemo() {
	newUser := user{
		name: "王五",
		age:  23,
	}
	sqlStr := "insert into user(name,age) values (?,?)"
	// the `Exec` method returns a `Result` type instead of a `Row`
	// we follow the same argument pattern to add query params
	result, err := db.Exec(sqlStr, newUser.name, newUser.age)
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	theID, err := result.LastInsertId() //新插入的id
	if err != nil {
		log.Fatalf("get LastInsertId falid. err:%v", err)
	}
	fmt.Printf("insert succ. the id is %d\n", theID)

}

// 更新数据
func updateRowDemo() {
	newUser := user{
		name: "王五",
		age:  28,
	}
	sqlStr := "update user set age=? where name = ?"
	result, err := db.Exec(sqlStr, newUser.age, newUser.name)
	if err != nil {
		log.Fatalf("update falid. err:%v", err)
	}

	// the `Result` type has special methods like `RowsAffected` which returns the
	// total number of affected rows reported by the database
	// In this case, it will tell us the number of rows that were inserted using
	// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get affected rows: %v", err)
	}
	// we can log how many rows were update
	fmt.Println("update succ, affected rows:%v", rowsAffected)

}

// 删除数据
func deleteRowDemo() {
	newUser := user{
		name: "张三",
	}
	sqlStr := "delete from user where name = ?"
	result, err := db.Exec(sqlStr, newUser.name)
	if err != nil {
		log.Fatalf("update falid. err:%v", err)
	}

	// the `Result` type has special methods like `RowsAffected` which returns the
	// total number of affected rows reported by the database
	// In this case, it will tell us the number of rows that were inserted using
	// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get affected rows: %v", err)
	}
	// we can log how many rows were delete
	fmt.Println("delete succ, affected rows:%v", rowsAffected)
}

//3.0 测试超时
func timeOutDemo() {
	// create a parent context
	ctx := context.Background()
	// create a context from the parent context with a 300ms timeout
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)
	// The context variable is passed to the `QueryContext` method as
	// the first argument
	// the provided number of seconds. We can use this to simulate a
	// slow query
	_, err := db.QueryContext(ctx, "select sleep(1)")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
}

//4.0 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() //开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		fmt.Printf("begin trans failed. err:%v\n", err)
		return
	}

	sqlStr := "Update user set age = 34 where id = ?"
	result, err := tx.Exec(sqlStr, 2)
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		fmt.Printf("update trans failed. err:%v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		log.Fatalf("could not get affected rows: %v", err)
		return
	}
	fmt.Println(rowsAffected)

	if rowsAffected == 1 {
		fmt.Println("only one user affect. trans commit")
		tx.Commit() //提交事务
	} else {
		fmt.Println("affect num err. trans rollback")
		tx.Rollback()
	}

	fmt.Println("exec trans end")
}

// 5.0 预处理操作示例
func prepareQueryDemo() {
	sqlStr := "select id,name,age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare err.")
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("user:%v\n", u)
	}
}

// 第一步:测试db是否ping通
func pingDemo() {
	// The `sql.Open` function opens a new `*sql.DB` instance. We specify the driver name
	// and the URI for our database.
	dsn := "xxxx:xxxx@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")
	defer db.Close()
}
