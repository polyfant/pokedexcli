package main

import (
	"fmt"
)

func main() {
	fmt.Println(getStartupMessage())
	cfg := NewConfig()
	startRepl(cfg)
}
