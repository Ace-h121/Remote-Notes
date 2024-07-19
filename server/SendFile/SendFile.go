package sendfile

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"os"
)

func SendFile(notespath string, path string) ([]byte, error) {
    f, err := os.Open(notespath + path)
    if err != nil {
	return nil, err
    }
    defer f.Close()

    gr, err := gzip.NewReader(f)
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(gr)

    buffer := new(bytes.Buffer)

    for scanner.Scan(){
	_, err := buffer.Write([]byte(scanner.Text()))
	if err != nil {
		return nil, err
	}
	
    }

    return buffer.Bytes(), nil 
}
