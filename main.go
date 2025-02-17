package main

import (
	"flag"
	"github.com/mbassini/zenwork/tasks"
	"log"

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

	//err = lists.New("Fifth List")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//found, err := lists.Find(3)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(found)
	//
	//err = tasks.New(model.Task{Name: "Second"}, 2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = tasks.Complete(2, 2)
	if err != nil {
		log.Fatal(err)
	}
}
