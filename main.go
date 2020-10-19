package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GLOBAL VARIABLE
var Meetings []Meeting

// STRUCT FOR DEFINING MEETING
type Meeting struct {
	// Id                 string `json:"Id "`
	IdMeeting          string `json:"IdMeeting"`
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
		if id == Meetings[i].IdMeeting {

			// fmt.Fprintf(w, Meetings[i].Title)
			json.NewEncoder(w).Encode(Meetings[i])
		}
	}

	// json.NewEncoder(w).Encode(Meetings)
}

func MeetingOperations(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kshitij:qwerty123@cluster0.kfjyr.mongodb.net/appointly?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	fmt.Println("CONNECTION TO MONGO DB IS SUCCESSFULL")

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
				// fmt.Println("sad")
				json.NewEncoder(w).Encode(Meetings[i])
			}
		}
		fmt.Println("THE ROUTE FOR PARTICIPANT CORRECT DONE")
	}
	if len(presented_query["start"]) != 0 && len(presented_query["end"]) != 0 {
		for i := 0; i < len(Meetings); i++ {
			// layout := "2006-01-02T15:04:05.000Z"
			fmt.Println(presented_query["start"][0])
			fmt.Println(presented_query["end"][0])
			fmt.Println(Meetings)
			if presented_query["start"][0] == "\""+Meetings[i].Start_Time+"\"" && presented_query["end"][0] == "\""+Meetings[i].End_Time+"\"" {
				// fmt.Println("BOOM BOOM")
				json.NewEncoder(w).Encode(Meetings[i])
			}

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
		log.Println(inMeet)
		appointlyDB := client.Database("appointlyDB")
		meetingsCollection := appointlyDB.Collection("meetings")
		// participantCollection := appointlyDB.Collection("participants")

		meetingResult, err := meetingsCollection.InsertOne(ctx, bson.D{
			{Key: "Id", Value: inMeet.IdMeeting},
			{Key: "Title", Value: inMeet.Title},
			{Key: "Participants", Value: inMeet.Participants},
			{Key: "Start_Time", Value: inMeet.Start_Time},
			{Key: "End_Time", Value: inMeet.End_Time},
			{Key: "Creation_Timestamp", Value: inMeet.Creation_Timestamp}})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(meetingResult.InsertedID)
		just_now := time.Now()
		var newMeeting Meeting = Meeting{IdMeeting: inMeet.IdMeeting, Title: inMeet.Title, Participants: inMeet.Participants, Start_Time: inMeet.Start_Time, End_Time: inMeet.End_Time, Creation_Timestamp: just_now.String()}
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

	Meetings = []Meeting{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kshitij:qwerty123@cluster0.kfjyr.mongodb.net/appointly?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	fmt.Println("CONNECTION TO MONGO DB IS SUCCESSFULL")

	appointlyDB := client.Database("appointlyDB")
	meetingsCollection := appointlyDB.Collection("meetings")
	// participantCollection := appointlyDB.Collection("participants")

	// meetingResult, err := meetingsCollection.InsertOne(ctx, bson.D{
	// 	{Key: "Id", Value: "4"},
	// 	{Key: "Title", Value: "TESTING"},
	// 	{Key: "Participants", Value: "KSHITIJ DHYAN"},
	// 	{Key: "Start_Time", Value: "2020-10-19T01:00:21.572903+05:30"},
	// 	{Key: "End_Time", Value: "2020-10-19T03:00:21.572903+05:30"},
	// 	{Key: "Creation_Timestamp", Value: "2020-10-19T00:20:21.572903+05:30"}})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(meetingResult.InsertedID)

	cursor, err := meetingsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var allMeetingsDB []bson.M
	if err = cursor.All(ctx, &allMeetingsDB); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(allMeetingsDB)
	for i := 0; i < len(allMeetingsDB); i++ {
		// fmt.Println(allMeetingsDB[i]["IdMeeting"])
		// fmt.Printf("var1 = %T\n", allMeetingsDB[i]["IdMeeting"])
		// fmt.Println(fmt.Sprint(allMeetingsDB[i]["Start_Time"]))
		var newMeeting Meeting = Meeting{IdMeeting: fmt.Sprint(allMeetingsDB[i]["Id"]), Title: fmt.Sprint(allMeetingsDB[i]["Title"]), Participants: fmt.Sprint(allMeetingsDB[i]["Participants"]), Start_Time: fmt.Sprint(allMeetingsDB[i]["Start_Time"]), End_Time: fmt.Sprint(allMeetingsDB[i]["End_Time"]), Creation_Timestamp: fmt.Sprint(allMeetingsDB[i]["Creation_Timestamp"])}
		Meetings = append(Meetings, newMeeting)
	}
	// fmt.Println(Meetings)

	handleRequests()
}
