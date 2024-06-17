package transfer

import (
	"log"
	"os"
)

func CleanFile(filename string){
	err := os.Remove(filename)
	if err != nil{
		log.Fatal(err)
	}
}

