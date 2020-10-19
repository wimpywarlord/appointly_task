package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// GLOBAL VARIABLE
var Meetings []Meeting

// STRUCT FOR DEFINING MEETING
type Meeting struct {
	Id                 string    `json:"Id"`
	Title              string    `json:"Title"`
	Participants       string    `json:"Participants"`
	Start_Time         time.Time `json:"Start_Time"`
	End_Time           time.Time `json:"End_Time"`
	Creation_Timestamp time.Time `json:"Creation_Timestamp"`
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

			// fmt.Fprintf(w, Meetings[i].Title)
			json.NewEncoder(w).Encode(Meetings[i])
		}
	}

	// json.NewEncoder(w).Encode(Meetings)
}

func MeetingOperations(w http.ResponseWriter, r *http.Request) {
	var presented_query = r.URL.Query()
	// fmt.Printf("var1 = %T\n", presented_query["participant"][0])
	// fmt.Println(presented_query["participant"])
	// fmt.Println(presented_query["start"][0])
	// fmt.Println(presented_query["end"][0])
	if len(presented_query["participant"]) != 0 {
		w.Header().Set("Content-Type", "application/json")

		for i := 0; i < len(Meetings); i++ {
			fmt.Println(Meetings[i].Participants)
			fmt.Println(presented_query["participant"][0])
			if presented_query["participant"][0] == "\""+Meetings[i].Participants+"\"" {
				fmt.Println("sad")
				json.NewEncoder(w).Encode(Meetings[i])
			}
		}
		fmt.Println("THE ROUTE FOR PARTICIPANT CORRECT DONE")
	}
	if len(presented_query["start"]) != 0 && len(presented_query["end"]) != 0 {
		for i := 0; i < len(Meetings); i++ {
			// layout := "2006-01-02T15:04:05.000Z"
			// fmt.Println(presented_query["start"][0])
			// fmt.Println(presented_query["end"][0])
			start_meeting_time, err := time.Parse(time.ANSIC, presented_query["start"][0])
			end_meeting_time, err := time.Parse(time.ANSIC, presented_query["end"][0])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(start_meeting_time)
			fmt.Println(end_meeting_time)
			// if presented_query["participant"][0] == Meetings[i].Participants {

			// 	// fmt.Fprintf(w, Meetings[i].Title)
			// 	json.NewEncoder(w).Encode(Meetings[i])
			// }
		}
		fmt.Println("ROUTE FOR RANGE MEETING CORRECTLY DONE")
	}
	if len(presented_query["participant"]) == 0 && len(presented_query["start"]) == 0 && len(presented_query["end"]) == 0 {
		// reqBody, _ := ioutil.ReadAll(r.Body)
		// fmt.Fprintf(w, "%+v", string(reqBody))
		// fmt.Println(reqBody)
		decoder := json.NewDecoder(r.Body)
		var inMeet Meeting
		w.Header().Set("Content-Type", "application/json")
		err := decoder.Decode(&inMeet)
		if err != nil {
			panic(err)
		}
		// log.Println(inMeet.Title)
		just_now := time.Now()
		var newMeeting Meeting = Meeting{Id: inMeet.Id, Title: inMeet.Title, Participants: inMeet.Participants, Start_Time: inMeet.Start_Time, End_Time: inMeet.End_Time, Creation_Timestamp: just_now}
		Meetings = append(Meetings, newMeeting)
		json.NewEncoder(w).Encode(inMeet)
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println(body)
		// var inMeet Meeting
		// err = json.Unmarshal(body, &inMeet)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println(inMeet.Id)
		fmt.Println("SCHEDULE MEETING ROUTE")
	}

}

func handleRequests() {
	http.HandleFunc("/meeting/", returnMeetingOfId)
	http.HandleFunc("/meetings", MeetingOperations)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	Meetings = []Meeting{
		Meeting{Id: "1", Title: "KSHITIJ IS GOING TO GET SELECTED", Participants: "ks@gmail.com", Start_Time: time.Now().UTC(), End_Time: time.Now().UTC(), Creation_Timestamp: time.Now().UTC()},
		Meeting{Id: "2", Title: "KSHITIJ IS BEST CANDIDATE", Participants: "dhyani@gmail.com", Start_Time: time.Now().UTC(), End_Time: time.Now().UTC(), Creation_Timestamp: time.Now().UTC()},
		Meeting{Id: "3", Title: "KSHITIJ WILL WORK HARD", Participants: "varun@gmail.com", Start_Time: time.Now().UTC(), End_Time: time.Now().UTC(), Creation_Timestamp: time.Now().UTC()},
	}
	handleRequests()
}
