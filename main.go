package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func processFile(filename string) {
	fmt.Println("Processing file:", filename)
	// Simulate processing by reading the file contents
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error reading file:", filename, "-", err)
		return
	}
	fmt.Println("File contents:", string(content))
}

func main() {
	var wg sync.WaitGroup

	filenames := []string{"file1.txt", "subdirectory/file2.txt", "file3.txt", "subdirectory/file4.txt", "file5.txt"}

	for _, filename := range filenames {
		// pass the file name as a channal <--
		wg.Add(1)
		//
		go func(filename string) { //this func to be subsribed to the above channal
			defer wg.Done()

			// Get the absolute path of the file
			absPath, err := filepath.Abs(filename)
			if err != nil {
				log.Println("Error getting absolute path of file:", filename, "-", err)
				return
			}

			// Check if the file exists
			_, err = os.Stat(absPath)
			if os.IsNotExist(err) {
				log.Println("File not found:", absPath)
				return
			}

			processFile(absPath)
		}(filename)
	}

	wg.Wait()
	fmt.Println("All files processed.")
}
