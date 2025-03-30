package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(programPackageName())
}

func programPackageName() int {
	log.Println("Hello World")
	return 0
}
