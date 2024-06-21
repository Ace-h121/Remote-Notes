package main

import (
	"fmt"
	"log"
	"os"

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

	switch method {

	case Send:
		sendMethod(args)
		os.Exit(0)

	case Receive:
		receiveMethod(args)
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
		
		fmt.Println(listingMethod(path))
		os.Exit(0)

	default:
		fmt.Println("Couldnt understand method, exiting program")
		os.Exit(1)
	}

}

func sendMethod(args []string){

	for _, arg := range args{
		content, err := encrypt.PrepareFile(arg)

		if err != nil {
			panic(err)
		}

		file := transfer.MakeFileStruct(content, arg)
		err = transfer.SendFile(file, "http://localhost:8090/send")

		if err != nil {
			panic(err)
		}

		

	}

}

func receiveMethod(args []string){
	for _, arg := range args{
		file, err := transfer.RecieveFile(arg, "http://localhost:8090/recieve")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		file.Content, err = decrypt.DecryptFile(file.Content)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		

		transfer.CreateFile(file.Name, file.Content)
	}

}

func listingMethod(arg string) string {
	list, err := transfer.ListFiles(arg, "http://localhost:8090/list")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		
	}
	return string(list)
}
