package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/cli"
	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/io"
	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	filePath := os.Args[1]
	cmds, err := io.ExtractCommandsFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	dispatcher := cli.NewDispatcher()

	for _, line := range cmds {
		parts := strings.Split(line, " ")
		cmd := parts[0]
		args := parts[1:]

		err := dispatcher.Handle(cmd, args)
		if err != nil {
			utils.PrintCapitalizedError(err)
		}
	}
}
