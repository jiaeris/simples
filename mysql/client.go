package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

//https://github.com/go-sql-driver/mysql
//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

func Conn() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true&loc=Local", "root", "root", "tps")
	Db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	fmt.Println(Db.Ping())
}

func Run() {
	dbRoot := &DbRoot{
		Name:"test",
		Account:"root",
		Password:"root",
	}
	dbRoot.Conn()
}

type DbRoot struct {
	Name       string
	Account    string
	Password   string
	ClientConn *sql.DB //数据库客户端
	DbConn     *sql.DB //数据库连接实例
}

func (dbRoot *DbRoot) CreateDb() {

}

func (dbRoot *DbRoot) CreateTable() {

}

func (dbRoot *DbRoot) ShowTables() {
	//dbRoot.DbConn.
}

func (dbRoot *DbRoot) Conn() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true&loc=Local",
		dbRoot.Account, dbRoot.Password, dbRoot.Name)
	dbRoot.DbConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = dbRoot.DbConn.Ping()
	fmt.Println(err)
}

