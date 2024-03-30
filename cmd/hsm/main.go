package main

import (
	"os"

	"github.com/HOKK-FINANCE-FZCO/crypto11"
)

func main() {
	hctx, err := crypto11.Configure(&crypto11.Config{
		Path:       os.Args[1],
		TokenLabel: os.Args[2],
		Pin:        os.Args[3],
	})
	if err != nil {
		panic(err)
	}

	hctx.Close()
}
