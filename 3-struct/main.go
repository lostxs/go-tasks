package main

import (
	"3-struct/bin"
	"3-struct/storage"
	"fmt"
)

func main() {
	storage := storage.NewStorage()
	bin := bin.NewBin("2", "new bin 2", true)
	storage.AddBin(*bin)
	bins := storage.GetBins()
	fmt.Println(bins)
}
