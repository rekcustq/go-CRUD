package main

import (
	"os"
	"net/http"

	"github.com/rekcustq/qlns/handler"
)

func main() {
	os.Setenv("PORT", "9876")
	port := os.Getenv("PORT")
	// route
	// user
	http.HandleFunc("/listAllUsers", handler.ListAllUsers)
	http.HandleFunc("/getUser", handler.GetUser)
	http.HandleFunc("/createUser", handler.CreateUser)
	http.HandleFunc("/updateUser", handler.UpdateUser)
	http.HandleFunc("/deleteUser", handler.DeleteUser)
	// shift
	http.HandleFunc("/listAllShifts", handler.ListAllShifts)
	http.HandleFunc("/getShift", handler.GetShift)
	http.HandleFunc("/createShift", handler.CreateShift)
	http.HandleFunc("/updateShift", handler.UpdateShift)
	http.HandleFunc("/deleteShift", handler.DeleteShift)
	http.ListenAndServe(":" + port, nil)
}