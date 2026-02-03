package handlers

import (
	"bufio"
	"log"
	"morse-server/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := os.Chdir(".")
	if err != nil {
		log.Println("error in changing director:", err)
	}
	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Println("error in reading file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fileText := scanner.Text()
	res := service.TextTypeSwitch(fileText)

	createdFile, err := os.Create(time.Now().UTC().String() + filepath.Ext(fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer createdFile.Close()

	_, err = createdFile.Write([]byte(res))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
