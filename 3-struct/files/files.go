package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(fileName string) *JsonDb {
	return &JsonDb{fileName: fileName}
}

func (db *JsonDb) Read() ([]byte, error) {
	if ext := filepath.Ext(db.fileName); ext != ".json" {
		return nil, errors.New("invalid file extension")
	}
	return os.ReadFile(db.fileName)
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	if _, err = file.Write(content); err != nil {
		fmt.Println(err)
		return
	}
}
