package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unhanded/mytoman/internal/myto_api"
	"github.com/unhanded/mytoman/internal/util"
)

func main() {
	var filepath string
	var dataDir string
	flag.StringVar(&filepath, "f", "", "the selected myto file")
	flag.StringVar(&dataDir, "d", "$HOME/.myto", "myto data directory")

	flag.CommandLine.Usage = func() {
		util.Banner()

		fmt.Print("\tUsage: mytom -f <file>\n\n")
		fmt.Print("\tExample: mytom -f ./message.myto\n\n")
	}

	flag.Parse()

	if filepath == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	res := myto_api.MytoExec(filepath, dataDir)

	fmt.Printf("%s\n", res)
}
