package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("invalid name")
	}

	return fmt.Sprintf("Hello, %s !", name), nil
}

func main() {
	fmt.Println(hello("foo"))
}
