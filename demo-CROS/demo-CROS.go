package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.CourseId == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("no course with id %d", ID), http.StatusNotFound)
	}
	switch r.Method {
	case http.MethodGet:
		courseJson, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(courseJson)
	case http.MethodPut:
		var updateCourse Course //Interface
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byteBody, &updateCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateCourse.CourseId != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updateCourse
		CourseList[listItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> coursesHandler called:", r.Method, r.URL.Path)
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

func enableCorsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Origin, Accept, Content-Type, Content-Length, Authorization, X-Cross, X-Requested-With, ngrok-skip-browser-warning")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// รวมทั้ง /course และ /course/{id}
	mux.Handle("/course", http.HandlerFunc(coursesHandler))
	mux.Handle("/course/", http.HandlerFunc(courseHandler))

	// ครอบ CORS middleware ให้รับ OPTIONS ได้ทุกเส้นทาง
	handlerWithCors := enableCorsMiddleware(mux)

	fmt.Println("Server is running at :5000")
	log.Fatal(http.ListenAndServe(":5000", handlerWithCors))
}
