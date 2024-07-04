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

	"Github.com/Ace-h121/RecieveFile"
	sendfile "Github.com/Ace-h121/SendFile"
)

type File struct{
	Content []byte `json:"content"`
	Name string `json:"name"`
}

type Path struct {
	Path string `json:"path"`
}

	var notespath string 

func main(){
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find Home Dir")
	}
	config, err := os.ReadFile(homedir + "/.config/Remote_Notes_Server")
	if err != nil{
		fmt.Print("Could not find/understand config file")
		log.Fatal(err)
	}

	notespath = string(config)
	notespath = strings.Trim(notespath, "\n")
	fmt.Printf("default notes directory: %s \n" , notespath)


	if err != nil{
		log.Fatal(err)
	} 

	router := http.NewServeMux()
	router.HandleFunc("/send", handleRecieve)
	router.HandleFunc("/recieve", handleSend)
	router.HandleFunc("/list", handleList)
	http.ListenAndServe(":8080", router)
}

func handleRecieve(w http.ResponseWriter, r *http.Request){

	decoder := json.NewDecoder(r.Body)
	
	var file File

	err := decoder.Decode(&file)
	
	if err != nil {
		panic(err)
	}
	log.Printf("Reciving file %s from %s", file.Name, r.URL.String())
	go func(){
		err = recievefile.WriteFile(notespath ,file.Name, file.Content)
		if err != nil{
			log.Print(err)
		}
	}()
}

func handleSend(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var path Path

	err := decoder.Decode(&path)
	if err != nil {
		panic(err)
	}

	log.Printf("Sending file %s to %s", path.Path, r.URL.String())
	content, err := sendfile.SendFile(notespath ,path.Path)
	if err != nil {
		log.Print(err)
		w.Write([]byte(fmt.Sprint(err)))
	}

	filename := filepath.Base(path.Path)
	filename = strings.Trim(filename, ".gz")

	

	data, err := json.Marshal(File{
		Content: content,
		Name: filename,
	})

	if err != nil{
		log.Print(err)
	}

	w.Write(data)


}

func handleList(w http.ResponseWriter, r *http.Request){
	req, err := io.ReadAll(r.Body)
	if err != nil{
		log.Print(err)
	}
	log.Printf("Got request to list files starting at %s, from client %s", req, r.URL.String())
	dir, err := os.ReadDir(notespath + string(req))
	if err != nil{
		log.Print(err)
	}
	
	var resp string

	ansiTeal := "\033[36m"

	ansiReset := "\033[0m"

	for _, file := range dir{

		if file.IsDir(){
		resp = resp + ansiTeal + file.Name() + ansiReset + " "
		} else {
		 resp = resp+ file.Name() + " "
		}
	}

	w.Write([]byte(resp))
	

}

