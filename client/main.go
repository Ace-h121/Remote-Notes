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
		fmt.Println("Couldnt understand message")
	}

}

func sendMethod(args []string){

	for _, arg := range args{
		content, err := encrypt.PrepareFile(arg)

		if err != nil {
			log.Fatal(err)
		}

		file := transfer.MakeFileStruct(content, arg)
		err = transfer.SendFile(file, "")

		if err != nil {
			log.Fatal(err)
		}

	}

}

func receiveMethod(args []string){
	for _, arg := range args{
		decrypt.DownloadFile(arg)
		transfer.CleanFile(arg)
	}

}
