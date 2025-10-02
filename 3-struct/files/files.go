package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	if ext := filepath.Ext(name); ext != ".json" {
		return nil, errors.New("invalid file extension")
	}
	return os.ReadFile(name)
}

func WriteFile(name string, content []byte) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	if _, err = file.Write(content); err != nil {
		fmt.Println(err)
		return
	}
}
