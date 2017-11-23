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
