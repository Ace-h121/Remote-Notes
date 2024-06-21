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
	Send = "-s"
	Receive = "-m"
)

func main(){
	method := os.Args[1]
	args := os.Args[2:]

	fmt.Println(method)

	switch method {

	case Send:
		sendMethod(args)

	case Receive:
		receiveMethod(args)

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
		file, err := transfer.RecieveFile(arg, "localhost")

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
