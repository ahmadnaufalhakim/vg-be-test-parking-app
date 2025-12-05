package main

import (
	"fmt"
	"os"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	cmds, err := internal.ExtractCommandsFromFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cmd := range cmds {
		fmt.Println(cmd)
	}
}
