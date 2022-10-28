package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story struct {
	Arcs map[string]Arc
	t    *template.Template
}

func makeTemplate(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func main() {
	jsonBytes, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("Unable to open story file.")
	}

	arcs := make(map[string]Arc)
	err = json.Unmarshal(jsonBytes, &arcs)
	if err != nil {
		log.Fatal("Unable to parse story file.")
	}

	cyoaTemplate := template.Must(template.ParseFiles("cyoa.html"))
	fmt.Println("Serving story on http://localhost:8080")
	http.ListenAndServe(":8080", Story{arcs, cyoaTemplate})

}

func (s Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimLeft(r.URL.Path, "/")

	if key == "" {
		key = "intro"
	}

	arc, ok := s.Arcs[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	s.t.Execute(w, arc)
}
