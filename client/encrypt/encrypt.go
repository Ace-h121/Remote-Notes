package encrypt

import (
	"os"
	"io"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"compress/gzip"
	"encoding/base64"
)

func openFile(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil{
		fmt.Printf("Got Error Opening file: %s", err.Error())
		panic(err)
	}
	return string(contents)
}

func encryptFile(contents string) string {
	block, err := aes.NewCipher([]byte("UcEqnUpzNoqYpb1O5kpormNFcpd7CNG0"))
	if err != nil {
		panic(err)
	}

	plainText := []byte(contents)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(ciphertext)
}

func createFile(contents string, filename string) {
	file, err := os.Create(filename + ".gz")
	if err != nil{
		fmt.Printf("Got Error Creating file %s", err.Error())
		panic(err)
	}
	gz := gzip.NewWriter(file)
	_, err = gz.Write([]byte(contents))
	
	if err != nil{
		fmt.Printf("Got Error Writing to File %s", err.Error())
		panic(err)
	}

	gz.Close()
	file.Close()

}

func PrepareFile(fileName string){
	createFile(encryptFile(openFile(fileName)), fileName)
}
