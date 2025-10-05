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

func (s *StorageWithDb) UpdateBin(updated *bin.Bin) {
	for i, b := range s.Bins {
		if b.ID == updated.ID {
			s.Bins[i] = *updated
			s.save()
			return
		}
	}
	s.AddBin(updated)
}

func (s *StorageWithDb) DeleteBin(id string) {
	var bins []bin.Bin
	for _, b := range s.Bins {
		if b.ID != id {
			bins = append(bins, b)
		}
	}
	s.Bins = bins
	s.save()
}

func (s *StorageWithDb) GetBins() []bin.Bin {
	var bins []bin.Bin
	bins = append(bins, s.Bins...)
	return bins
}

func (s *StorageWithDb) save() {
	data, err := json.Marshal(s.Storage)
	if err != nil {
		fmt.Println(err)
	}
	s.db.Write(data)
}
