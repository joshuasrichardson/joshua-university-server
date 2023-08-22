package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GenerateCoursesResponse struct {
	Courses []Course `json:"courses"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func generateCoursesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	subject := r.URL.Query().Get("subject")

	log.Printf(subject)

	courses := []Course{
		{
			ID:          "CS 142",
			Name:        "Introduction to Computer Science",
			Description: "Introduction to Computer Science",
		},
		{
			ID:          "CS 224",
			Name:        "Computer Systems",
			Description: "Introduction to Computer Science",
		},
		{
			ID:          "CS 235",
			Name:        "Data Structures",
			Description: "Introduction to Computer Science",
		},
		{
			ID:          "CS 240",
			Name:        "Advanced Programming Concepts",
			Description: "Introduction to Computer Science",
		},
	}

	response := GenerateCoursesResponse{
		Courses: courses,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/generate-courses", generateCoursesHandler)
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// [
//                 {
//                   id: "CS 142",
//                   name: "Introduction to Computer Science",
//                   description: "Introduction to Computer Science",
//                 },
//                 {
//                   id: "CS 224",
//                   name: "Computer Systems",
//                   description: "Introduction to Computer Science",
//                 },
//                 {
//                   id: "CS 235",
//                   name: "Data Structures",
//                   description: "Introduction to Computer Science",
//                 },
//                 {
//                   id: "CS 240",
//                   name: "Advanced Programming Concepts",
//                   description: "Introduction to Computer Science",
//                 },
//               ]