package main

import (
	"days"
	"os"
)

func main() {
	//NOTE: This will error out if not set :p
	args := os.Args
	switch args[1] {
	case "1":
		days.One()
	case "2":
		days.Two()
	case "3":
		days.Three()
	}
}
