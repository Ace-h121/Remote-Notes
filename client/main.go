package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"Github.com/Ace-h121/decrypt"
	"Github.com/Ace-h121/encrypt"
	"Github.com/Ace-h121/transfer"
)


const(
	Send = "upload"
	Receive = "download"
	Help = "help"
	List = "list"
)

func main(){
	method := os.Args[1]
	args := os.Args[2:]

	homedir, err := os.UserHomeDir()
	if err != nil{
		log.Fatal("Could not find home dir")
		os.Exit(1)
	}

	config, err := os.Open(homedir + "/.config/Remote_Notes")
	if err != nil{
		log.Fatal("Could not find config file")
		os.Exit(1)
	}
	defer config.Close()
	
	content, err := io.ReadAll(config)

	key := string(content)

	key, ipaddr, _ := strings.Cut(key, ";")



	switch method {

	case Send:
		sendMethod(args, key, ipaddr)
		os.Exit(0)

	case Receive:
		receiveMethod(args, key, ipaddr)
		os.Exit(0)

	case Help:
		fmt.Println(`Usage:
	upload [file1 file2 ...]   : Encrypt and upload specified files
	download [file1 file2 ...] : Download and decrypt specified files
	list [path (optional)]	   : List all files in the dir, if no given dir lists root
	help                       : Display this help message

Commands:
	upload       Uploads and encrypts the specified files to the server.
	download     Downloads and decrypts the specified files from the server.
	list		 Lists the given directory
	help         Displays this help message.
		`)
		os.Exit(0)

	case List:
		var path string
		if len(args) < 1{
			path = ""
		}else {
			path = args[0]
		}
		
		fmt.Println(listingMethod(path, ipaddr))
		os.Exit(0)

	default:
		fmt.Println("Couldnt understand method, exiting program")
		os.Exit(1)
	}

}

func sendMethod(args []string, key string, ipaddr string){

	for _, arg := range args{
		content, err := encrypt.PrepareFile(arg, key)

		if err != nil {
			panic(err)
		}

		file := transfer.MakeFileStruct(content, arg)
		err = transfer.SendFile(file, ipaddr + "/send")
		if err != nil {
			panic(err)
		}

	}

}

func receiveMethod(args []string, key string, ipaddr string){
	for _, arg := range args{
		file, err := transfer.RecieveFile(arg, ipaddr + "/recieve")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		file.Content, err = decrypt.DecryptFile(file.Content, key)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		

		transfer.CreateFile(file.Name, file.Content)
	}

}

func listingMethod(arg string, ipaddr string) string {
	list, err := transfer.ListFiles(arg, ipaddr +"/list")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		
	}
	return string(list)
}
