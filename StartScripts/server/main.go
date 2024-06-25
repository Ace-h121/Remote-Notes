package main

import (
	"log"
	"os"
)

func main(){
	if len(os.Args)< 2{
		log.Fatal("Please pass in the full desired directory for notes to be servered")
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Couldnt locate homedir")
	}
	notespath := os.Args[1]
	
	file, err := os.Create(homedir + "/.config/Remote_Notes_Server")
	defer file.Close()

	file.Write([]byte(notespath))

}
