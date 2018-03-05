package main

import (
	"fmt"

	"github.com/makki0205/config"
)

func main() {
	fmt.Println(config.Env("dialect"))

	fmt.Println(config.Env("foo"))
}
