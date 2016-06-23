package user

import (
	"log"

	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

type User struct {
	Name       string   `json:"name"`
	Password   string   `json:"password"`
	Namespaces []string `json:"namespaces"`
	Cpus       int      `json:"cpus"`
	Memory     int      `json:"memory"`
	CpusUse    int      `json:"cpususe"`
	MemoryUse  int      `json:"memoryuse"`
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

func (user *User) GetNamespacesAll(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	dbuser.Name = user.Name

	db, err := dbuser.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	res, err := dbuser.QueryAll(db)
	if err != nil || len(res) == 0 {
		return err
	}

	user.Namespaces = res
	return nil
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

func (user *User) CreateUser(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)
	userinfo.Name = user.Name
	userinfo.Password = user.Password
	userinfo.Cpus = user.Cpus
	userinfo.Mem = user.Memory

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = userinfo.Insert(db)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) DeleteUser(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)
	userinfo.Name = user.Name

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = userinfo.Delete(db)
	if err != nil {
		log.Println("DeleteUser error :", err)
		return err
	}

	return nil
}

func (user *User) GetUser(dbconf *sqlop.MysqlCon) error {
	userinfo := new(sqlop.UserInfo)
	userinfo.Name = user.Name

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	defer db.Close()

	err = userinfo.Query(db)
	if err != nil {
		return err
	}

	user.Password = userinfo.Password
	user.Cpus = userinfo.Cpus
	user.Memory = userinfo.Mem

	return nil
}

func (user *User) GetAllInfo(dbconf *sqlop.MysqlCon) error {
	dbuser := new(sqlop.User)
	userinfo := new(sqlop.UserInfo)
	userinfo.Name = user.Name
	dbuser.Name = user.Name

	db, err := userinfo.Connect(dbconf)
	if err != nil {
		return err
	}

	user.Namespaces, err = dbuser.Query(db)
	if err != nil {
		log.Println(err)
		return err
	}

	err = userinfo.Query(db)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = userinfo.Password
	user.Cpus = userinfo.Cpus
	user.Memory = userinfo.Mem

	return nil
}
