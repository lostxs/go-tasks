package bin

type BinList struct {
	Bins []Bin
}

func NewBinList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}
