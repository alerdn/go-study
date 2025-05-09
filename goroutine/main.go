package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func writer(file *os.File, data <-chan string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for strData := range data {
		_, err := file.WriteString(strData)
		if err != nil {
			panic(err)
		}
	}
}

func startWorkers(qtdWorkers int, file *os.File, data <-chan string) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(qtdWorkers)

	for range qtdWorkers {
		go writer(file, data, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
	}()
}

func main() {
	log.Println("Starting...")
	file, err := os.Create("./file/file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(chan string)

	qtdWorkers := runtime.NumCPU() * 3
	startWorkers(qtdWorkers, file, data)

	for range 10000000 {
		data <- time.Now().String() + "\n"
	}

	time.Sleep(1 * time.Second)
	log.Println("Finished writing to file")
}
