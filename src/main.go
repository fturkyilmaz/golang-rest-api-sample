package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"src/model"
)

const (
	endpoint = "https://random-data-api.com/api/users/random_user?size=10"
)

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllUsers")

	var responseObject []model.User

	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	errJson := json.Unmarshal(responseData, &responseObject)

	if errJson != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(responseObject); i++ {
		fmt.Println(responseObject)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response ID:", responseObject[0].ID)

	jData, err := json.Marshal(responseObject)

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func handleHttpServer() {
	fmt.Println("Http Server Started Listen Port :8080")

	// Set routing rules
	http.HandleFunc("/", defaultRoute)
	http.HandleFunc("/users", getAllUsers)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleHttpServer()
}
