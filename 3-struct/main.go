package main

import (
	"fmt"
	"time"
)

type bin struct {
	id        string
	name      string
	private   bool
	createdAt time.Time
}

func newBin(id, name string, private bool) *bin {
	return &bin{
		id:        id,
		name:      name,
		private:   private,
		createdAt: time.Now(),
	}
}

type binList struct {
	bins []bin
}

func (bl *binList) printBinList() {
	fmt.Println(bl)
}

func (bl *binList) addBin(bin *bin) {
	bl.bins = append(bl.bins, *bin)
}

func newBinList() *binList {
	return &binList{
		bins: []bin{},
	}
}

func main() {
	fmt.Println("Bins prog")
	binList := newBinList()
	bin := newBin("1", "first bin", false)
	binList.addBin(bin)
	binList.printBinList()
}
