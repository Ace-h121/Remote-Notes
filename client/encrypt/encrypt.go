package encrypt

import (
	"os"
	"fmt"
	"crypto/aes"
	"compress/gzip"
)

func openFile(file string) string {
	pwd, err := os.Getwd()
	if err != nil{
		fmt.Printf("Got Error Opening file: %s", err.Error())
		panic(err)
	}
	contents, err := os.ReadFile(pwd + file)
	if err != nil{
		fmt.Printf("Got Error Opening file: %e", err.Error())
		panic(err)
	}
	return string(contents)
}

func encryptFile(contents string) string {
	block, err := aes.NewCipher([]byte("Team5401!!"))
	if err != nil{
		fmt.Printf("Got Error Encrypting contents: %s", err.Error())
	}
	out := make([]byte, len(contents))
	block.Encrypt(out, []byte(contents))
	return string(out) 
}

func createFile(contents string) {
	pwd, err := os.Getwd()
	if err != nil{
		fmt.Printf("Got Error Opening file: %s", err.Error())
		panic(err)
	}
	file, err := os.Create(pwd + "temp.gz")
	if err != nil{
		fmt.Printf("Got Error Creating file %s", err.Error())
		panic(err)
	}
	gz := gzip.NewWriter(file)
	_, err := gz.Write([]byte(contents))
	
	if err != nil{
		fmt.Printf("Got Error Writing to File %s", err.Error())
		panic(err)
	}

	file.Close()

}


