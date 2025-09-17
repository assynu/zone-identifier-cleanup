package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: zone-identifier-cleanup <directory>")
	}

	dir := os.Args[1]

	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Fatalf("Error: directory '%s' does not exist", dir)
	}
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if !info.IsDir() {
		log.Fatalf("Error: '%s' is not a directory", dir)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go removeFilesInDirectory(dir, &wg)

	wg.Wait()
}

func removeFilesInDirectory(directory string, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(directory, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go removeFilesInDirectory(entryPath, wg)
		} else if strings.Contains(entry.Name(), "Zone.Identifier") {
			log.Printf("Delete %s", entryPath)
			os.Remove(entryPath)
		}
	}
}
