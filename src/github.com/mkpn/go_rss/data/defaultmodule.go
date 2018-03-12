package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mkpn/go_rss/secret"
)

// https://github.com/golang/go/wiki/SQLDrivers
func Insert() {
	mysqlpass := secret.Password()
	db, err := sql.Open("mysql", "root:"+mysqlpass+"@/rss")
	print(err)
	print(db)
	db.Exec(`insert into test (memo) values (?) `, "fuga")
}
