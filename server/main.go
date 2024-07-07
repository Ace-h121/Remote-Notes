package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	recievefile "Github.com/Ace-h121/RecieveFile"
	sendfile "Github.com/Ace-h121/SendFile"
)

//struct for sending files
type File struct {
	Content []byte `json:"content"`
	Name    string `json:"name"`
}

// struct for path, eaiser and safe to encode as json
type Path struct {
	Path string `json:"path"`
}

var notespath string

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find Home Dir")
	}

	//look for config
	config, err := os.ReadFile(homedir + "/.config/Remote_Notes_Server")
	if err != nil {
		fmt.Print("Could not find/understand config file\n")
		notespath = "/app/notes/"
		err = nil
	} else {
		notespath = string(config)
		notespath = strings.Trim(notespath, "\n")
	}

	//parse the config

	fmt.Printf("default notes directory: %s \n", notespath)

	router := http.NewServeMux()
	router.HandleFunc("/send", handleRecieve)
	router.HandleFunc("/recieve", handleSend)
	router.HandleFunc("/list", handleList)
	http.ListenAndServe(":8080", router)
}

func handleRecieve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	decoder := json.NewDecoder(r.Body)

	var file File

	err := decoder.Decode(&file)

	if err != nil {
		panic(err)
	}
	log.Printf("Reciving file %s from %s", file.Name, r.RemoteAddr)
	go func() {
		err = recievefile.WriteFile(notespath, file.Name, file.Content)
		if err != nil {
			log.Print(err)
		}
	}()

}

func handleSend(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	decoder := json.NewDecoder(r.Body)

	var path Path

	err := decoder.Decode(&path)
	if err != nil {
		panic(err)
	}

	log.Printf("Sending file %s to %s", path.Path, r.RemoteAddr)
	content, err := sendfile.SendFile(notespath, path.Path)
	if err != nil {
		log.Print(err)
		w.Write([]byte(fmt.Sprint(err)))
	}

	filename := filepath.Base(path.Path)
	filename = strings.Trim(filename, ".gz")

	data, err := json.Marshal(File{
		Content: content,
		Name:    filename,
	})

	log.Println(filename)

	if err != nil {
		log.Print(err)
	}

	w.Write(data)

}

func handleList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Got request to list files starting at %s, from client %s", req, r.RemoteAddr)
	dir, err := os.ReadDir(notespath + string(req))
	if err != nil {
		log.Print(err)
	}

	var resp string

	ansiTeal := "\033[36m"

	ansiReset := "\033[0m"

	for _, file := range dir {

		if file.IsDir() {
			resp = resp + ansiTeal + file.Name() + ansiReset + " "
		} else {
			resp = resp + file.Name() + " "
		}
	}

	w.Write([]byte(resp))

}
