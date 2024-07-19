package recievefile

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile(notespath string, filename string, contents string) error {
	fmt.Printf("Writing file, %s, to %s \n", filename+".gz", notespath)
	os.MkdirAll(filepath.Dir(notespath+filename), 0700)

	file, err := os.Create(notespath + filename + ".gz")
	if err != nil {
		return err
	}


	gz := gzip.NewWriter(file)
	_, err = gz.Write([]byte(contents))

	if err != nil {
		return err
	}

	gz.Close()
	file.Close()

	return nil
}
