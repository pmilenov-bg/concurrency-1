package main

import (
	"fmt"
	"sync"
	"time"
)

func welcomeMessage(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}

func sendToChannel(filename string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	ch <- fmt.Sprintf("%s is been processed", filename)
}

func printChannel(ch chan string, wg *sync.WaitGroup) {
	for data := range ch {
		fmt.Println(data)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go welcomeMessage("hello", &wg)
	fmt.Println("go routine started")

	ch := make(chan string)
	filenames := []string{"file1.txt", "subdirectory/file2.txt", "file3.txt", "subdirectory/file4.txt", "file5.txt"}

	// Increment wait group count to match the number of files being processed
	wg.Add(len(filenames))

	// Launch goroutines to process files and send data to the channel
	for _, filename := range filenames {
		go sendToChannel(filename, ch, &wg)
	}

	fmt.Println("lets wait")

	// Use a separate goroutine for printing the channel data since it will block
	go printChannel(ch, &wg)

	wg.Wait()
	fmt.Println("last stage")
	close(ch)
}
