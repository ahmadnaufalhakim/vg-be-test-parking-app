package utils

import (
	"fmt"
	"strings"
)

func PrintCapitalizedError(err error) {
	errMsg := err.Error()
	errMsg = strings.ToUpper(errMsg[:1]) + errMsg[1:]
	fmt.Println(errMsg)
}
