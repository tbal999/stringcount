package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type stringcount struct {
	String             string `json:"String"`
	Numberofcharacters int    `json:"Numberofcharacters"`
}

func getPort() string {
	p := os.Getenv("PORT")
	fmt.Println(p)
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func countString(w http.ResponseWriter, r *http.Request) {
	S := Stringcount{}
	S.String = mux.Vars(r)["string"]
	S.Numberofcharacters = len(S.String)
	json.NewEncoder(w).Encode(S)
}

func help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "String count API - Counts the number of characters in any string!\n")
	fmt.Fprintf(w, "Commands:\n")
	fmt.Fprintf(w, `1) /count/{string} - counts the number of characters in the string i.e /count/hi would be {"String":"hi","Number of characters":2}`+"\n")
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "String count API - Counts the number of characters in any string!")
}

func main() {
	port := getPort()
	fmt.Println("API has started.")
	fmt.Println("Running on port " + port)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", help)
	router.HandleFunc("/count/{string}", countString).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))

}
