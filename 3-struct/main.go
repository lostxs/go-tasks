package main

import (
	"3-struct/api"
	"3-struct/bin"
	"3-struct/config"
	"3-struct/files"
	"3-struct/storage"
	"flag"
	"fmt"
)

func main() {
	create := flag.Bool("create", false, "Create new bin")
	update := flag.Bool("update", false, "Update bin")
	delete := flag.Bool("delete", false, "Delete bin")
	get := flag.Bool("get", false, "Get bin")
	list := flag.Bool("list", false, "List bins")

	file := flag.String("file", "", "Path to file with bins data")
	name := flag.String("name", "", "Bin name")
	id := flag.String("id", "", "Bin id")
	flag.Parse()

	cfg := config.NewConfig()
	api := api.NewApi(cfg)
	storage := storage.NewStorage(files.NewJsonDb("bins.json"))

	switch {
	case *create:
		resp, err := api.CreateBin(*file, *name)
		if err != nil {
			fmt.Println(err)
			return
		}
		bin := bin.NewBin(resp.Metadata.ID, *name, resp.Metadata.Private)
		storage.AddBin(bin)
		fmt.Printf("Bin создан: %s", bin.ID)
	case *update:
		_, err := api.UpdateBin(*file, *id)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := api.GetBin(*id)
		if err != nil {
			fmt.Println(err)
			return
		}
		updatedBin := bin.NewBin(resp.Metadata.ID, *name, resp.Metadata.Private)
		storage.UpdateBin(updatedBin)
		fmt.Printf("Bin обновлен: %s", updatedBin.ID)
	case *delete:
		resp, err := api.DeleteBin(*id)
		if err != nil {
			fmt.Println(err)
			return
		}
		storage.DeleteBin(resp.Metadata.ID)
		fmt.Printf("Bin удален: %s", resp.Metadata.ID)
	case *get:
		resp, err := api.GetBin(*id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Bin: %+v\n", resp.Record)
	case *list:
		bins := storage.GetBins()
		for _, bin := range bins {
			bin.Print()
		}
	default:
		flag.Usage()
	}
}
