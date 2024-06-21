package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)



func DecryptFile(content []byte) ([]byte, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(string(content))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte("UcEqnUpzNoqYpb1O5kpormNFcpd7CNG0"))
	if err != nil {	
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, err
}



