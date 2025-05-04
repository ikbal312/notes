// /*
// // synchronization in go

// method 1 ; WaitGroup
// method 2 : Channel
// method 3 : Mutex and RWMutex

// */
// /*

// // synchronization using WaitGroup

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			fmt.Println(n)
// 		}(i)
// 	}
// 	wg.Wait() //Wait blocks until the WaitGroup counter is zero.
// 	fmt.Println("all work done")
// }
// */

// /*
// Synchronization using Channel

// channel are use to pass data from one go routine to another go routine
// by giving access to a single go routine at a time.
// */

// /*
// 	// channel to channel communication
// 	ch := make(chan string, 1) // creating channel
// 	go func() {
// 		for _, word := range []string{"hello", "world"} {
// 			ch <- word // sending word to channel
// 		}
// 		defer close(ch) // closing channel and channel should only close by sender
// 	}()

// 	// receiving data
// 	for word := range ch { // acts like a <- ch
// 		fmt.Println(word)
// 	}

// */

// /* select statement

// inCh1 := make(chan string, 2)
// // inCh2 := make(chan string, 2)
// go func() {
// 	for _, word := range []string{"word1", "word2"} {
// 		inCh1 <- word
// 		close(inCh1)
// 	}

// 	}()
// // go func() {
// // 	for _, word := range []string{"word1", "word2"} {
// 	// 		inCh2 <- word
// // 	}
// // 	close(inCh2)
// // }()

// for {
// 	select {
// 	case v := <-inCh1:
// 		go fmt.Println("received(intCh1):", v)
// 		// case v := <-inCh2:
// 			// 	go fmt.Println("received(intCh2):", v)
// 		default:
// 		fmt.Println("channel was empty")
// 	}

// }
// */

// /*
// 	Channel as Event Signal

// package main

// import (

// 	"fmt"
// 	"sync"

// )

// func main() {

// 	in1 := make(chan string)
// 	in2 := make(chan string)
// 	exit := make(chan struct{})
// 	wg := &sync.WaitGroup{}

// 	wg.Add(1)

// 	go printWords(in1, in2, exit, wg)

// 	in1 <- "hello"
// 	in2 <- "word"
// 	close(exit)
// 	wg.Wait()

// }

// 	func printWords(in1, in2 chan string, exit chan struct{}, wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-exit:
// 				fmt.Println("exiting")
// 				return
// 			case str := <-in1:
// 				fmt.Println("in1: ", str)
// 			case str := <-in2:
// 				fmt.Println("in2: ", str)
// 			}
// 		}
// 	}
// */

// /*
// Mutexes
// mutex is a synchronization primitive also known as lock
// */
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type sum struct {
// 	mu  sync.Mutex
// 	sum int
// }

// func (s *sum) get() int {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	return s.sum
// }

// func (s *sum) add(n int) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.sum += n
// }

// func main() {

// 	mySum := &sum{}
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(x int) {
// 			defer wg.Done()
// 			mySum.add(x)
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println("final sum:", mySum.get())
// }
