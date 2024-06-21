package main

import (
	"encoding/json"
	"log"
	"net/http"

)

type File struct{
	Content []byte `json:"content"`
	Name string `json:"name"`
}

type Path struct {
	Path string `json:"path"`
}

func main(){
	router := http.NewServeMux()
	router.HandleFunc("POST /send", handleRecieve)


	http.ListenAndServe("localhost:8090", router)
}

func handleRecieve(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	
	var file File

	err := decoder.Decode(&file)
	
	if err != nil {
		panic(err)
	}

	log.Print(file)
}


