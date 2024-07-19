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
	Content string `json:"content"`
	Name string `json:"name"`
}

type Path struct {
	Path string `json:"path"`
}

func MakeFileStruct(content string, name string) File{
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


	if err != nil {
		log.Fatal("Requested file does not exist")
	}

	return newFile, nil 
}

func CreateFile(filename string, content string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	file.Write([]byte(content))
	return nil
}

func ListFiles(path string, url string ) ([]byte, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(path)))
	if err != nil{
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil{
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
