package filefirewall

import (
	"log"
	"os"
)

const FILENAME = "logs.txt"

func LogDebug(contents string) {
	createFile()
	f := openFile()
	writeToFile(f, contents)
	defer f.Close()
}

func createFile() {
	_, err := os.Create(FILENAME)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func openFile() *os.File {
	f, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return f
}

func writeToFile(f *os.File, contents string) {
	_, err := f.WriteString("[DEBUG] " + contents + "\n")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
