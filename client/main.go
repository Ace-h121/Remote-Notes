package main

import (
	"os"

	"Github.com/Ace-h121/encrypt"
)

func main(){
	args := os.Args[1:]
	for _, arg := range args{
		encrypt.PrepareFile(arg)
	}
}
