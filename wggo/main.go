package main

import "fmt"
import "sync"

func printNums(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	for i := range ch {
		fmt.Println("Printing ", i)
	}
}

func generateNums(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()

	for i:=0; i<30; i++ {
		fmt.Println("Generating ", i)
		ch <- i
	}
}

func main() {
	var wg sync.WaitGroup

	channel := make(chan int)

	wg.Add(2)
	go printNums(&wg, channel)
	generateNums(&wg, channel) //both of these run concurrently
	defer wg.Wait()
	close(channel)

}