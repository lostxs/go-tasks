package bins

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

func NewBin(id, name string, private bool) *bin {
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

func (bl *binList) PrintBinList() {
	fmt.Println(bl)
}

func (bl *binList) AddBin(bin *bin) {
	bl.bins = append(bl.bins, *bin)
}

func NewBinList() *binList {
	return &binList{
		bins: []bin{},
	}
}
