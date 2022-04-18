package l337C0d3

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func handleNums(wg *sync.WaitGroup, mergeChan chan<- string, numChan chan<- bool, charChan <-chan bool, i int) {
	defer wg.Done()
	for j := 1; j <= i; j++ {
		<-charChan
		s := strconv.Itoa(j)
		mergeChan <- s
		numChan <- true
	}
	// Unlock char handler (the last one to end)
	<-charChan
}

func handleChars(wg *sync.WaitGroup, mergeChan chan<- string, numChan <-chan bool, charChan chan<- bool, i int) {
	defer wg.Done()
	// Being the last handler to end, close the merge one to stop its loop
	defer close(mergeChan)
	for j := 0; j < i; j++ {
		<-numChan
		s := string(rune(j + 'A'))
		mergeChan <- s
		charChan <- true
	}
}

func merger(wg *sync.WaitGroup, c <-chan string, builder *strings.Builder) {
	defer wg.Done()
	for s := range c {
		builder.WriteString(s)
	}
}

func alternateRoutines(i int) (string, error) {
	if i < 0 || i > 'z'-'A' {
		return "", fmt.Errorf("i must be between 0 and %d", 'z'-'A')
	}
	builder := strings.Builder{}
	mergeChan := make(chan string)
	numChan := make(chan bool)
	charChan := make(chan bool)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go merger(wg, mergeChan, &builder)
	go handleNums(wg, mergeChan, numChan, charChan, i)
	go handleChars(wg, mergeChan, numChan, charChan, i)

	// Unlock num handler (the first one to start)
	charChan <- true

	wg.Wait()
	return builder.String(), nil
}
