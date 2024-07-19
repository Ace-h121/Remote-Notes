package encrypt

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

func openFile(file string) (string, error) {
	contents, err := os.ReadFile(file)
	if err != nil{
		return "handle your errors pls pls pls", nil
	}
	return string(contents), nil
}

func encryptFile(contents string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plainText := []byte(contents)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}


	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(ciphertext), nil
}

func createFile(contents string, filename string) error {
	file, err := os.Create(filename + ".gz")
	if err != nil{
		return err
	}
	gz := gzip.NewWriter(file)
	_, err = gz.Write([]byte(contents))
	
	if err != nil{
		return err
	}

	gz.Close()
	file.Close()

	return nil
}

func PrepareFile(fileName string, key string) (string, error){
	content, err := openFile(fileName)

	if err != nil{
		return "", err
	}

	encryptedContent, err := encryptFile(content, key)

	
	if err != nil{
		return "", err
	}

	return encryptedContent, nil

}
