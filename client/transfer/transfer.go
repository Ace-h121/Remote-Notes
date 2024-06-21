package transfer

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func CleanFile(filename string){
	err := os.Remove(filename)
	if err != nil{
		log.Fatal(err)
	}
}


type File struct {
	Content []byte `json:"content"`
	Name string
}

type Path struct {
	Path string `json:"path"`
}

func MakeFileStruct(content []byte, name string) File{
	return File {
		Content: content,
		Name: name,
	}
}

func SendFile(file File, url string) error{
	jsonstr, err := json.Marshal(file)
	if err != nil{
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	if err != nil{
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	defer req.Body.Close()
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	
	return nil
}

func RecieveFile(path string, url string) (File, error){
	message := Path{
		Path: path,
	}

	jsonstr, err := json.Marshal(message)

	if err != nil{
		return File{}, err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonstr))
	if err != nil{
		return File{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	defer req.Body.Close()

	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil{
		return File{}, err
	}

	body, err := io.ReadAll(resp.Body)

	var newFile File

	err = json.Unmarshal(body, &newFile)

	return newFile, nil 
}
