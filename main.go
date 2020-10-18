package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GLOBAL VARIABLE
var Meetings []Meeting

// STRUCT FOR DEFINING MEETING
type Meeting struct {
	Id                 string `json:"Id"`
	Title              string `json:"Title"`
	Participants       string `json:"Participants"`
	Start_Time         string `json:"Start_Time"`
	End_Time           string `json:"End_Time"`
	Creation_Timestamp string `json:"Creation_Timestamp"`
}

// STRUCT FOR DEFINING PARTICIPANT
type Participant struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

// END POINT FOR "GET A MEETING USING ID"
func returnMeetingOfId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint For: Listing Meeting of id X")
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	// fmt.Println(id[2])

	for i := 0; i < len(Meetings); i++ {
		if id == Meetings[i].Id {
			fmt.Fprintf(w, Meetings[i].Title)
		}
	}

	// json.NewEncoder(w).Encode(Meetings)
}

func MeetingOperations(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	var presented_query = r.URL.Query()
	// fmt.Printf("var1 = %T\n", presented_query["participant"][0])
	fmt.Println(presented_query["participant"])
	// fmt.Println(presented_query["start"][0])
	// fmt.Println(presented_query["end"][0])
	if len(presented_query["participant"]) != 0 {
		fmt.Println("THE ROUTE FOR PARTICIAPNT CORRECT DOEN")
	}
	if len(presented_query["start"]) != 0 && len(presented_query["end"]) != 0 {
		fmt.Println("ROUTE FOR RANGE MEETING CORRECTLY DONE")
	}
	if len(presented_query["participant"]) == 0 && len(presented_query["start"]) == 0 && len(presented_query["end"]) == 0 {
		fmt.Println("SCHEDULE MEETING ROUTE")
	}

}

func handleRequests() {
	http.HandleFunc("/meeting/", returnMeetingOfId)
	http.HandleFunc("/meetings/", MeetingOperations)
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
