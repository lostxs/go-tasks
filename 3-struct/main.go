package main

import (
	"3-struct/bin"
	"3-struct/files"
	"3-struct/storage"
	"fmt"
)

func main() {
	storage := storage.NewStorage(files.NewJsonDb("bins.json"))
	bin := bin.NewBin("2", "new bin 2", true)
	storage.AddBin(bin)
	bins := storage.GetBins()
	fmt.Println(bins)
}
