package main

import (
	"github.com/itsemre/find-my/pkg/apple"
)

func main() {
	_, err := apple.Login("", "", true)
	if err != nil {
		panic(err)
	}
}
