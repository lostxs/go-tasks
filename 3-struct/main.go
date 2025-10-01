package main

import (
	"3-struct/bins"
	"fmt"
)

func main() {
	fmt.Println("Bins prog")
	binList := bins.NewBinList()
	bin := bins.NewBin("1", "first bin", false)
	binList.AddBin(bin)
	binList.PrintBinList()
}
