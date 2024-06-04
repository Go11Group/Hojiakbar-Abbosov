package main

import (
	"fmt"
	"sync"
	"time"
)

func sendData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var input int
		fmt.Print("Iltimos, ma'lumot kiriting (0 - dasturni to'xtatish): ")
		fmt.Scanln(&input)

		if input == 0 {
			close(ch) 
			return
		}

		ch <- input
	}
}

func displayData(ch chan int, wg *sync.WaitGroup) {
 	defer wg.Done()
 
 	for input := range ch {
  		fmt.Println("Qabul qilingan ma'lumot:", input)
 	}
 		fmt.Println("Channel yopilgan, dastur to'xtatiladi.")
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go sendData(ch, &wg)      
	go displayData(ch, &wg)    

	wg.Wait()

	
}