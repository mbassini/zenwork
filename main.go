package main

import (
	"fmt"

	"github.com/mbassini/zenwork/internal/storage"
)

func main() {
	db := storage.InitDB()
	fmt.Printf("%v", db)
}
