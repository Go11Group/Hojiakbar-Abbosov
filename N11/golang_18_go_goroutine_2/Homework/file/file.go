package file

import (
	"encoding/json"
	"my_module/book"
	"os"
	"sync"

)

func WriteBooks(ch <-chan []book.Book, done chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Create("books.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for books := range ch {
		err := encoder.Encode(books)
		if err != nil {
			panic(err)
		}
	}
	done <- struct{}{}
}

func ReadBooks(ch chan<- []book.Book, done chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open("books.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for {
		var books []book.Book
		err := decoder.Decode(&books)
		if err != nil {
			break 
		}
		ch <- books
	}
	done <- struct{}{}
}
