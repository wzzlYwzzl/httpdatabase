package users

import (
	"fmt"
	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

type Users struct {
	Name       string   `json:"name"`
	Namespaces []string `json:"namespaces"`
}

func (user Users) JudgeExist(dbconf *sqlop.MysqlCon) (bool, error) {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	db, err := sqlop.Connect(dbconf)
	if err != nil {
		return false, err
	}

	defer db.Close()

	res, err := sqlop.Query(db, dbuser)
	if err != nil || len(res) == 0 {
		return false, err
	}

	return true, nil
}

func (user *Users) GetNamespaces(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	res, err := sqlop.Query(db, dbuser)
	if err != nil || len(res) == 0 {
		return err
	}

	user.Namespaces = res
	return nil
}
