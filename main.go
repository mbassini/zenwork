package main

import (
	"flag"
	//"fmt"
	"log"
	//"time"

	//"github.com/mbassini/zenwork/model"
	"github.com/mbassini/zenwork/storage"
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Println(arg)
	}

	err := storage.LoadLists()
	if err != nil {
		log.Fatal(err)
	}
}
