package filefirewall

import (
	"fmt"
	"os"
)

const (
	NOT_ALLOWED_ERROR_MESSAGE = "You are not allowed to read this file unfortunately."
)

func Fail() {
	fmt.Fprintf(os.Stderr, NOT_ALLOWED_ERROR_MESSAGE)
	os.Exit(2)
}

func Succeed() {
	os.Exit(0)
}
