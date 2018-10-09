package main

import (
	"log"
	"os"
)

func main() {
	log.Println("smart monkeys is starting...")
}

func writeConfiguration(destination string, content string) {
	file, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteSize, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}

	log.Printf("wrote %d bytes on %s", byteSize, destination)
}
