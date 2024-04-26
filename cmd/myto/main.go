package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unhanded/mytomake/internal/myto_api"
	"github.com/unhanded/mytomake/internal/util"
)

func main() {
	var filepath string
	var dataDir string
	flag.StringVar(&filepath, "f", "", "the selected myto file")
	flag.StringVar(&dataDir, "d", "$HOME/.myto", "myto data directory")

	flag.CommandLine.Usage = func() {
		util.Banner()

		fmt.Print("\tUsage: myto -f <file>\n\n")
		fmt.Print("\tExample: myto -f ./message.myto\n\n")
	}

	flag.Parse()

	if filepath == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	res := myto_api.MytoExec(filepath, dataDir)

	fmt.Printf("%s\n", res)
}
