package main

import (
	"net/http"
	"html/template"
	"fmt"
)
type PageData struct{
	Result string
	Text string
	Banner string
}
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Path not found", 404)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}

}
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusAccepted)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "Text cannot be empty", http.StatusBadRequest)
		return
	}
	if banner == "" {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	if banner != "shadow" && banner != "thinkertoy" && banner != "standard" {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}

	char, err := ValidateInput(text)
	if err != nil {
		http.Error(w, fmt.Sprintf("%c is not a valid character", char), http.StatusBadRequest)
		return
	}
	charMap, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	result := GenerateArt(text, charMap)

	data := PageData{
		Result: result,
		Text: text,
		Banner: banner,
	}
	err = tmpl.ExecuteTemplate(w, "result.html", data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-switch" {
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Failed to parse form", http.StatusAccepted)
	// 	return
	// }
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	if text == "" {
		http.Error(w, "Text cannot be empty", http.StatusBadRequest)
		return
	}
	if banner == "" {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	if banner != "shadow" && banner != "thinkertoy" && banner != "standard" {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}

	char, err := ValidateInput(text)
	if err != nil {
		http.Error(w, fmt.Sprintf("%c is not a valid character", char), http.StatusBadRequest)
		return
	}
	charMap, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	result := GenerateArt(text, charMap)

	data := PageData{
		Result: result,
		Text: text,
		Banner: banner,
	}
	err = tmpl.ExecuteTemplate(w, "result.html", data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}

}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.HandleFunc("/ascii-art-switch", switchHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}