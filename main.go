package main

import (
	"os"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "c":
		Client()
	case "s":
		Server()
	}
}
