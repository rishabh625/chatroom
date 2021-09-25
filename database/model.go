package database

import (
	"crypto/md5"
	"fmt"
	"log"
)

type LoginRequest struct {
	Username string
	Password string
}

//Login Authenticates User Credentials from db, gives response authentication success, admin status
func Login(request *LoginRequest) (bool, bool) {
	db = GetConnection()
	if request.Username == "" || request.Password == "" {
		return false, false
	}
	query := "Select admin,username from users where username = $1 and password = $2 and enabled = $3"
	password := fmt.Sprintf("%x", md5.Sum([]byte(request.Password)))
	var admin bool
	var username string
	rows, err := db.Query(query, request.Username, password, false)
	if err != nil {
		return false, false
	}
	if rows.Next() {
		rows.Scan(&admin, &username)
		return true, admin
	}
	return false, false
}

type Message struct {
	Data     []byte
	Room     string
	Username string
}

//AddData Adds Chat History into Database, Returns Success,failure or prints Error as Message
func AddData(request Message) bool {
	db = GetConnection()
	// jsondata, err := json.Marshal(request.Data)
	// if err != nil {
	// 	log.Println(err)
	// 	return false
	// }
	insForm, err := db.Prepare("INSERT INTO chathistory (data,room,username) values ($1,$2,$3)")
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = insForm.Exec(string(request.Data), request.Room, request.Username)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
