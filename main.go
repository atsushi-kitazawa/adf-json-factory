package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	initWeb()
}

func initWeb() {
	port := "8080"
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/dataset", handleDataset)
	http.HandleFunc("/output", handleOutput)
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func handleOutput(w http.ResponseWriter, r *http.Request) {
	//var p DatasetPostData
	//if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	//	panic(err)
	//}

	dataset := NewDataset("test-dataset", "ArmtemplateStorageLinkedService", "test-storage", "emp.txt", "input")
	jsonData, err := json.Marshal(dataset)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("dataset.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(jsonData)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	if err := t.Execute(w, struct {
		Title   string
		Message string
		Time    time.Time
	}{
		Title:   "テストページ",
		Message: "こんにちは！",
		Time:    time.Now(),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func handleDataset(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/dataset.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	if err := t.Execute(w, struct{}{}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}
