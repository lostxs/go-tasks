package storage

import (
	"3-struct/bin"
	"encoding/json"
	"fmt"
)

type Db interface {
	Read() ([]byte, error)
	Write(content []byte)
}

type Storage struct {
	Bins []bin.Bin `json:"bins"`
}

type StorageWithDb struct {
	Storage
	db Db
}

func NewStorage(db Db) *StorageWithDb {
	file, err := db.Read()
	if err != nil {
		return &StorageWithDb{
			db: db,
			Storage: Storage{
				Bins: []bin.Bin{},
			},
		}
	}
	var storage Storage
	if err = json.Unmarshal(file, &storage); err != nil {
		fmt.Println(err)
	}
	return &StorageWithDb{
		db:      db,
		Storage: storage,
	}
}

func (s *StorageWithDb) AddBin(bin *bin.Bin) {
	s.Bins = append(s.Bins, *bin)
	s.save()
}

func (s *StorageWithDb) GetBins() []bin.Bin {
	var bins []bin.Bin
	bins = append(bins, s.Bins...)
	return bins
}

func (s *StorageWithDb) save() {
	data, err := s.toJsonBytes()
	if err != nil {
		fmt.Println(err)
	}
	s.db.Write(data)
}

func (s *Storage) toJsonBytes() ([]byte, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
