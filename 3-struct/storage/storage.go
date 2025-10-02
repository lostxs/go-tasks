package storage

import (
	"3-struct/bin"
	"3-struct/files"
	"encoding/json"
	"fmt"
)

type Storage struct {
	Bins []bin.Bin `json:"bins"`
}

func NewStorage() *Storage {
	file, err := files.ReadFile("bins.json")
	if err != nil {
		return &Storage{
			Bins: []bin.Bin{},
		}
	}
	var storage Storage
	if err = json.Unmarshal(file, &storage); err != nil {
		fmt.Println(err)
	}
	return &storage
}

func (s *Storage) AddBin(bin bin.Bin) {
	s.Bins = append(s.Bins, bin)
	s.save()
}

func (s *Storage) GetBins() []bin.Bin {
	var bins []bin.Bin
	bins = append(bins, s.Bins...)
	return bins
}

func (s *Storage) toJSONBytes() ([]byte, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (s *Storage) save() {
	data, err := s.toJSONBytes()
	if err != nil {
		fmt.Println(err)
	}
	files.WriteFile("bins.json", data)
}
