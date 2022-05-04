package mysql_custom

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/wsqyouth/tech_note/ddd_go/aggregate"
	"log"
	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
	// customer 用来存储customers
	// customer *mysql.Collection
}

// mysqlCustomer is an internal type that is used to store a CustomerAggregate
// we make an internal struct for this to avoid coupling this mongo implementation to the customeraggregate.
// Mongo uses bson so we add tags for that
type mysqlCustomer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromCustomer(c aggregate.Customer) mysqlCustomer {
	return mysqlCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m mysqlCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c

}

// Create a new mysqldb repository
func New(ctx context.Context, connectionString string) (*MysqlRepository, error) {
	// connectionString:= "coopers:2019Youth@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", connectionString) //这里暂时这是ip,port,密码外部传入,表名也外部写入
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{
		db: db,
	}, nil
}

func (mr *MysqlRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	fmt.Printf("get: %v\n", id)

	sqlStr := "select name from user where id = ? limit 1"

	// `QueryRow` always returns a single row from the database
	row := mr.db.QueryRow(sqlStr, 5) //固定值测试

	var c mysqlCustomer
	if err := row.Scan(&c.Name); err != nil {
		log.Fatalf("could not scan row: %v", err)
		return aggregate.Customer{}, err
	}
	fmt.Printf("user: %+v\n", c)

	// Convert to aggregate
	return c.ToAggregate(), nil
}

func (mr *MysqlRepository) Add(c aggregate.Customer) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	internal := NewFromCustomer(c)
	fmt.Printf("insert: %v,%v\n", internal.ID, internal.Name)
	// 这里暂时不改测试表了,仅插入name
	sqlStr := "insert into user(name) values (?)"
	result, err := mr.db.Exec(sqlStr, internal.Name)
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
		return err
	}
	theID, err := result.LastInsertId() //新插入的id
	if err != nil {
		log.Fatalf("get LastInsertId falid. err:%v", err)
		return err
	}
	fmt.Printf("insert succ. the id is %d\n", theID)

	return nil
}

func (mr *MysqlRepository) Update(c aggregate.Customer) error {
	panic("to implement")
}
