package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GLOBAL VARIABLE
var Meetings []Meeting

type Meeting struct {
	Id                 string `json:"Id"`
	Title              string `json:"Title"`
	Participants       string `json:"Participants"`
	Start_Time         string `json:"Start_Time"`
	End_Time           string `json:"End_Time"`
	Creation_Timestamp string `json:"Creation_Timestamp"`
}

type Participant struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// END POINT FOR "GET A MEETING USING ID"
func returnMeetingOfId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint For: Listing Meeting of id X")
	json.NewEncoder(w).Encode(Meetings)
}

func handleRequests() {
	http.HandleFunc("/", returnMeetingOfId)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	Meetings = []Meeting{
		Meeting{Id: "1", Title: "KSHITIJ IS GOING TO GET SELECTED", Participants: "2", Start_Time: "123", End_Time: "456", Creation_Timestamp: "789"},
		Meeting{Id: "2", Title: "KSHITIJ IS BEST CANDIDATE", Participants: "2", Start_Time: "123", End_Time: "456", Creation_Timestamp: "789"},
		Meeting{Id: "3", Title: "KSHITIJ WILL WORK HARD", Participants: "2", Start_Time: "123", End_Time: "456", Creation_Timestamp: "789"},
	}
	handleRequests()
}
