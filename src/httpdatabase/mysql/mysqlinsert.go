package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Name      string
	Namespace string
}

func Insert(db *sql.DB, user User) error {
	stmt, err := db.Prepare("INSERT INTO zjw(id, name, namespace) VALUE(null, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	stmt.Exec(user.Name, user.Namespace)
	return nil
}
