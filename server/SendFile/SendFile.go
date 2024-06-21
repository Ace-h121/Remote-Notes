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
	buffer.Write([]byte(scanner.Text()))
    }

    return buffer.Bytes(), nil 
}
