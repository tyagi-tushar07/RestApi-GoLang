package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Age    string `json:"Age"`
	Gender string `json:"Gender"`
}

type alluser []user

var users = alluser{
	{
		ID:     "1",
		Name:   "Tushar Tyagi",
		Age:    "22",
		Gender: "Male",
	},
}

func createuser(w http.ResponseWriter, r *http.Request) {
	var newuser user

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter the data in given format")
	}

	json.Unmarshal(reqbody, &newuser)
	users = append(users, newuser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newuser)

}

func getoneuser(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]

	for _, singleuser := range users {
		if singleuser.ID == userid {
			json.NewEncoder(w).Encode(singleuser)
		}

	}

}

func getalluser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)

}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]

	for i, singleuser := range users {

		if singleuser.ID == userid {
			users = append(users[:i], users[:i+1]...)
			fmt.Fprintf(w, "The User is Successfully deleted", userid)

		}

	}

}

func updateuser(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]

	var updateuser user

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please Enter write Pattern")
	}

	json.Unmarshal(reqbody, &updateuser)

	for i, singleuser := range users {

		if singleuser.ID == userid {
			singleuser.ID = updateuser.ID
			singleuser.Name = updateuser.Name
			singleuser.Gender = updateuser.Gender
			singleuser.Age = updateuser.Age
			users = append(users[:i], singleuser)
			json.NewEncoder(w).Encode(singleuser)

		}

	}

}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Tushar Tyagi & Go !!!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", start)
	router.HandleFunc("/user", createuser).Methods("POST")
	router.HandleFunc("/user/{id}", getoneuser).Methods("GET")
	router.HandleFunc("/users", getalluser).Methods("GET")
	router.HandleFunc("/users/delete/{id}", deleteuser).Methods("DELETE")
	router.HandleFunc("/users/update/{id}", updateuser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}
