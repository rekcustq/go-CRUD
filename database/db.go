package database

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int			`json:"id,omitempty"`
	Name string 	`json:"name"`
	Gender string	`json:"gender"`
	DOB string		`json:"date_of_birth"`
	Address string	`json:"address"`
	Phone string	`json:"phone"`
	Mail string		`json:"mail"`
	Position string	`json:"position"`
}

type Shift struct {
	Id int			`json:"id,omitempty"`
	Name string		`json:"name,omitempty"`
	UserId int	 	`json:"userid,omitempty"`
	StartTime string`json:"start_time"`
	EndTime string	`json:"end_time"`
}

func DbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "hades123"
    dbName := "qlns"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        log.Println(err)
    }
    return db
}
