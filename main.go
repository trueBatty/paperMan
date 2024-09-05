package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	currentDir, _ := os.Getwd()
	parentDir := filepath.Dir(currentDir)
	filename := os.Args[1]
	for _, dir := range os.Args[2:] {
		path := filepath.Join(parentDir, dir)
		os.MkdirAll(path, os.ModePerm)

		err := os.Symlink(filepath.Join(currentDir, filename), filepath.Join(path, filename))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
