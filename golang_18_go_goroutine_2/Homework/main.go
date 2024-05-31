package main

import (
	"fmt"
	"sync"

	"my_module/book"
	"my_module/file"
)

func main() {
	ch := make(chan []book.Book, 5)
    done := make(chan struct{})
    wg := sync.WaitGroup{}

    wg.Add(2)
    go file.WriteBooks(ch, done, &wg)
    go file.ReadBooks(ch, done, &wg)

    for _, b := range book.Books {
        ch <- []book.Book{b}
    }
    close(ch) 

    go func() {
        wg.Wait()
        close(done) 
    }()

    for {
        select {
        case <-done:
            return 
        case books := <-ch:
            for _, b := range books {
                fmt.Println(b)
            }
        }
    }
}
