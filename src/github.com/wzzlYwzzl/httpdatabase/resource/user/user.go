package user

import (
	"log"

	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

type User struct {
	Name       string   `json:"name"`
	Namespaces []string `json:"namespaces"`
}

func (user User) JudgeExist(dbconf *sqlop.MysqlCon) (bool, error) {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return false, err
	}

	defer db.Close()

	res, err := dbuser.Query(db)
	if err != nil || len(res) == 0 {
		log.Printf("return false")
		return false, err
	}

	return true, nil
}

func (user *User) GetNamespaces(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	res, err := dbuser.Query(db)
	if err != nil || len(res) == 0 {
		return err
	}

	user.Namespaces = res
	return nil
}

func (user *User) CreateNamespace(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name
	dbuser.Namespace = user.Namespaces[0]

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = dbuser.Insert(db)
	if err != nil {
		return err
	}

	return nil
}
