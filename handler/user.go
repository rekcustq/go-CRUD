package handler

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/rekcustq/qlns/database"

	_ "github.com/go-sql-driver/mysql"
)

func parseDOB(str string) string {
	date := strings.Split(str, "-")
	return date[2] + "-" + date[1] +  "-" + date[0]
}

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	allUsers, err := db.Query("SELECT * FROM user ORDER BY id")
    if err != nil {
        log.Println(err)
	}
	user := database.User{}
	result := []database.User{}
	for allUsers.Next() {
		var id int
        var name, gender, dob, address, phone, mail, position string
        err = allUsers.Scan(&id, &name, &gender, &dob, &address, &phone, &mail, &position)
        if err != nil {
            log.Println(err)
		}
		user.Id = id
		user.Name = name
		user.Gender = gender
		user.DOB = parseDOB(dob)
		user.Address = address
		user.Phone = phone
		user.Mail = mail
		user.Position = position
		result = append(result, user)
	}
	jsRes, _ := json.MarshalIndent(result, "", "    ")
	log.Println("Show all users")
	fmt.Fprintf(w, string(jsRes))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	id := r.URL.Query().Get("id")
	userInfo, err := db.Query("SELECT * FROM user WHERE id=?", id)
    if err != nil {
        log.Println(err)
	}
	user := database.User{}
	for userInfo.Next() {
        var id int
        var name, gender, dob, address, phone, mail, position string
        err = userInfo.Scan(&id, &name, &gender, &dob, &address, &phone, &mail, &position)
        if err != nil {
            log.Println(err)
		}
		user.Id = id
		user.Name = name
		user.Gender = gender
		user.DOB = parseDOB(dob)
		user.Address = address
		user.Phone = phone
		user.Mail = mail
		user.Position = position
	}
	jsRes, _ := json.MarshalIndent(user, "", "    ")
	log.Println("Get user", user.Id)
	fmt.Fprintf(w, string(jsRes))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	if r.Method == "POST" {
		user := database.User{}
        user.Name = r.FormValue("name")
		user.Gender = r.FormValue("gender")
		user.DOB = parseDOB(r.FormValue("dob"))
        user.Address = r.FormValue("address")
		user.Phone = r.FormValue("phone")
		user.Mail = r.FormValue("mail")
		user.Position = r.FormValue("position")
        addUser, err := db.Prepare("INSERT INTO User(name, gender, dob, address, phone, mail, position) VALUES(?,?,?,?,?,?,?)")
        if err != nil {
            log.Println(err)
        }
		_, err = addUser.Exec(user.Name, user.Gender, user.DOB, user.Address, user.Phone, user.Mail, user.Position)
		if err != nil {
			log.Println(err)
		}
		jsRes, _ := json.MarshalIndent(user, "", "    ")
		log.Println("Create user")
        fmt.Fprintf(w, string(jsRes))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
    if r.Method == "PUT" {
		user := database.User{}
        user.Name = r.FormValue("name")
		user.Gender = r.FormValue("gender")
		user.DOB = parseDOB(r.FormValue("dob"))
        user.Address = r.FormValue("address")
		user.Phone = r.FormValue("phone")
		user.Mail = r.FormValue("mail")
		user.Position = r.FormValue("position")
        user.Id, _ = strconv.Atoi(r.FormValue("id"))
        editUser, err := db.Prepare("UPDATE User SET name=?, gender=?, dob=?, address=?, phone=?, mail=?, position=? WHERE id=?")
        if err != nil {
            log.Println(err)
        }
		_, err = editUser.Exec(user.Name, user.Gender, user.DOB, user.Address, user.Phone, user.Mail, user.Position, user.Id)
		if err != nil {
			log.Println(err)
		}
		jsRes, _ := json.MarshalIndent(user, "", "    ")
		log.Println("Update user", user.Id)
        fmt.Fprintf(w, string(jsRes))
    }
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		delForm, err := db.Prepare("DELETE FROM User WHERE id=?")
		if err != nil {
			log.Println(err)
		}
		delForm.Exec(id)
		log.Println("Delete user", id)
		fmt.Fprintf(w, "User " + id + " deleted.")
	}
}