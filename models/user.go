package models

import (
	"log"
	"github.com/qowns8/ideaweb/utils"
)

var (
	db = utils.NewRDB()
)

type User struct {
	Id int
	Email string
	Pwd string
	Access_key string
	Name string
}

func GoogleAccessUpdate(token string, name string) {
	var id int
	e1 := db.Db.QueryRow("SELECT COUNT(id) FROM user WHERE email = ?", name).Scan(&id)
	if e1 != nil {
		log.Println(e1.Error())
		return
	}
	if id != 0 {
		db.Db.Exec(tokenUpdate, token, id)
	} else {
		db.Db.Exec(insert, name, "", token, name)
	}
}

func (user *User) Get() bool {
	if user.Email == "" && user.Pwd == "" {
		access_key := user.Access_key
		if access_key == "" {
			log.Println("login fail")
			return false
		}
		var update_time string
		err := db.Db.QueryRow(getInfoWithKey , access_key).Scan(&user.Id, &user.Email, &user.Pwd, &user.Access_key, &user.Name, &update_time)
		if err != nil {
			log.Println("access fail ")
			return false
		}
		log.Println("access success :  ")
		return true
	} else {
		err := db.Db.QueryRow(getInfoWithEmail, user.Email, user.Pwd).Scan(&user)
		if err != nil {
			log.Println("login fail")
			return false
		}
		log.Println("login : " + user.Email + "   " + user.Access_key)
		return true
	}
}