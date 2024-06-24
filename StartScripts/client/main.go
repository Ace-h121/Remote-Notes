package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/thanhpk/randstr"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatal("Do not have enough args, please provide the server address")
	}
	OS := getOS()
	log.Println("Finding OS type")
	log.Println(OS)
	makeConfig(OS)
}


func getOS() string{
	if runtime.GOOS == "windows"{
		return "windows"
	}
	if runtime.GOOS == "linux"{
		return "linux"
	}
	return ""
}

func makeConfig(OS string){
	if OS == "linux"{
		homedir, err :=  os.UserHomeDir()
		if err != nil {
			log.Fatal("Had error reading homedir")
			os.Exit(1)
		}
		configPath := filepath.Join(homedir + "/.config")
		_, err = os.Create(configPath+"/Remote_Notes")

		if err != nil{
			log.Fatal("Couldnt Create config file")		
		}

		key := randstr.Hex(16)

		ipaddr := os.Args[1]

		output := key + ";" + ipaddr

		err = os.WriteFile(configPath+"/Remote_Notes", []byte(output), 0644)
		
		if err != nil{
			log.Fatal("Couldnt Write File")
		}

	}
}
