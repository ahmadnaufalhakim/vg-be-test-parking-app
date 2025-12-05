package utils

import (
	"fmt"
	"strings"
)

func PrintCapitalizedError(err error) {
	errMsg := err.Error()
	if len(errMsg) == 0 {
		fmt.Println()
		return
	}

	errMsg = strings.ToUpper(errMsg[:1]) + errMsg[1:]
	fmt.Println(errMsg)
}
