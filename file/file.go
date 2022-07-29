package file

import (
	"io"
	"net/http"
	"os"
)

func Download(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// Creates a clean copy of the file
// Client is the responsible of closing the file
func Truncate(f *os.File) (*os.File, error) {
	err := os.Remove(f.Name())
	if err != nil {
		return nil, err
	}

	return os.Create(f.Name())
}

// Read all content and close the file
func Read(filename string) (string, *os.File, error) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0600)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", nil, err
	}

	return string(content), f, nil
}
