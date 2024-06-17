package main

import (
	"fmt"
	"os"

	"Github.com/Ace-h121/decrypt"
	"Github.com/Ace-h121/encrypt"
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
		encrypt.PrepareFile(arg)
	}

}

func receiveMethod(args []string){
	for _, arg := range args{
		decrypt.DownloadFile(arg)
	}

}
