package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"
	"os/exec"
	"strings"
)

func unzip(filename string) []byte {
	out, err := exec.Command("zcat", filename).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func decryptFile(content []byte) []byte {
	ciphertext, err := base64.RawStdEncoding.DecodeString(string(content))
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher([]byte("UcEqnUpzNoqYpb1O5kpormNFcpd7CNG0"))
	if err != nil {
		log.Fatal(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		log.Fatal(err)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

func createFile(filename string, content []byte){
	file, err := os.Create(strings.Trim(filename, ".gz"))
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.Write(content)
	
}



func DownloadFile(filename string){
	createFile(filename, decryptFile(unzip(filename)))
}
