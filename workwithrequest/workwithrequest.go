package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	CourseId   int    `json:"id"`
	CourseName string `json:"name"`
	Price      int    `json:"price"`
	Instructor string `json:"instructor"`
}

var CourseList []Course

func init() {
	CourseJson := `[
		{
			"id":1,
			"name":"Python",
			"price":2590,
			"instructor":"BorntoDev"
		},
		{
			"id":2,
			"name":"JavaScript",
			"price":0,
			"instructor":"BorntoDev"
		},
		{
			"id":3,
			"name":"SQL",
			"price":0,
			"instructor":"BorntoDev"
		}
	]`

	err := json.Unmarshal([]byte(CourseJson), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, course := range CourseList {
		if highestID < course.CourseId {
			highestID = course.CourseId
		}
	}
	return highestID + 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJson, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(courseJson)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newCourse.CourseId != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.CourseId = getNextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}

}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
