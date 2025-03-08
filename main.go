package main

import (
	"os"

	"github.com/mbassini/zenwork/cmd"
	"github.com/mbassini/zenwork/internal/storage"
)

func main() {
	storage.InitDB()
	cmd.Run(os.Args[1:])
}
