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

func parseTime(str string) string {
	tmp := strings.Split(str, " ")
	date := strings.Split(tmp[0], "-")
	return date[2] + "-" + date[1] +  "-" + date[0] + " " + tmp[1]
}

func ListAllShifts(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	allShifts, err := db.Query("SELECT s.id, u.name, u.id, s.starttime, s.endtime FROM user u INNER JOIN shift s ON s.userid = u.id ORDER BY u.id")
    if err != nil {
        log.Println(err)
	}
	shift := database.Shift{}
	result := []database.Shift{}
	for allShifts.Next() {
		var id, userid int
        var name, startTime, endTime string
        err = allShifts.Scan(&id, &name, &userid, &startTime, &endTime)
        if err != nil {
            log.Println(err)
		}
		shift.Id = id
		shift.Name = name
		shift.UserId = userid
		shift.StartTime = parseTime(startTime)
		shift.EndTime = parseTime(endTime)
		result = append(result, shift)
	}
	jsRes, _ := json.MarshalIndent(result, "", "    ")
	log.Println("Show all shifts")
	fmt.Fprintf(w, string(jsRes))
}

func GetShift(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	id := r.URL.Query().Get("id")
	shiftInfo, err := db.Query("SELECT s.id, u.name, u.id, s.starttime, s.endtime FROM user u INNER JOIN shift s ON s.userid = u.id WHERE s.id=?", id)
    if err != nil {
        log.Println(err)
	}
	shift := database.Shift{}
	for shiftInfo.Next() {
		var id, userid int
        var name, startTime, endTime string
        err = shiftInfo.Scan(&id, &name, &userid, &startTime, &endTime)
        if err != nil {
            log.Println(err)
		}
		shift.Id = id
		shift.Name = name
		shift.UserId = userid
		shift.StartTime = parseTime(startTime)
		shift.EndTime = parseTime(endTime)
	}
	jsRes, _ := json.MarshalIndent(shift, "", "    ")
	log.Println("Get shift", shift.Id)
	fmt.Fprintf(w, string(jsRes))
}

func CreateShift(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	if r.Method == "POST" {
		shift := database.Shift{}
        shift.UserId, _ = strconv.Atoi(r.FormValue("userid"))
		shift.StartTime = r.FormValue("starttime")
		shift.EndTime = r.FormValue("endtime")
        addShift, err := db.Prepare("INSERT INTO shift(userid, starttime, endtime) VALUES(?,?,?)")
        if err != nil {
            log.Println(err)
        }
		_, err = addShift.Exec(shift.UserId, shift.StartTime, shift.EndTime)
		if err != nil {
			log.Println(err)
		}
		jsRes, _ := json.MarshalIndent(shift, "", "    ")
		log.Println("Create shift")
        fmt.Fprintf(w, string(jsRes))
	}
}

func UpdateShift(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	if r.Method == "PUT" {
		shift := database.Shift{}
        shift.UserId, _ = strconv.Atoi(r.FormValue("userid"))
		shift.StartTime = r.FormValue("starttime")
		shift.EndTime = r.FormValue("endtime")
        shift.Id, _ = strconv.Atoi(r.FormValue("id"))
        addShift, err := db.Prepare("UPDATE User SET userid=?, starttime=?, endtime=? WHERE id=?")
        if err != nil {
            log.Println(err)
        }
		_, err = addShift.Exec(shift.UserId, shift.StartTime, shift.EndTime, shift.Id)
		if err != nil {
			log.Println(err)
		}
		jsRes, _ := json.MarshalIndent(shift, "", "    ")
		log.Println("Create shift")
        fmt.Fprintf(w, string(jsRes))
	}
}

func DeleteShift(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	defer db.Close()
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		delForm, err := db.Prepare("DELETE FROM shift WHERE id=?")
		if err != nil {
			log.Println(err)
		}
		delForm.Exec(id)
		log.Println("Delete shift", id)
		fmt.Fprintf(w, "Shift " + id + " deleted.")
	}
}