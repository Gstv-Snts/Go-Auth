package database

import (
	"database/sql"
	"fmt"
)

type User struct{
	Login string `json:"login"`
	Password string `json:"password"`
}
type UserRow struct{
	Id int
	Login string
	Password string 
}

func FindUser(db *sql.DB, login string)(UserRow,bool) {
	var userRow UserRow
	rows, err := db.Query(fmt.Sprintf(`select * from user where login = "%v"`, login))
	if err != nil {
		fmt.Println("#####################################")
		panic(err.Error())
	}
	rows.Next()
	if err := rows.Scan(&userRow.Id,&userRow.Login,&userRow.Password);err != nil{
		fmt.Println(err.Error())
	}
	rows.Close()
	return userRow, rows.Next()
}