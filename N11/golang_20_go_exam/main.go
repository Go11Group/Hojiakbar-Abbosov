package main

import (
	"encoding/json"
	"fmt"
	"os"
	"my_module/tasks"
)

type Input struct {
	Slc1  []int
	Text2 string
	Num   int
}

type Answer struct {
	Sum    int
	Titled string
	Unum   int
}

func main() {
	c := []int{23, 45, 6, 2, 34, 65, 546, 27, 54}
	c1 := "heLLo EorLd"
	c2 := -46
	c3 := make(chan int)

	a := Input{
		Slc1:  c,
		Text2: c1,
		Num:   c2,
	}

	fmt.Println(a)
	
	res1, err := tasks.SumOflice(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in SumOfSlice: %v\n", err)
		return
	}

	res2, err := tasks.TextToTitle(c1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in TextToTitle: %v\n", err)
		return
	}

	go func() {
		c3 <- tasks.Abs(c2)
	}()

	res3 := <-c3

	ans := Answer{
		Sum:    res1,
		Titled: res2,
		Unum:   res3,
	}

	f1, err := os.Create("flies/output.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}
	defer f1.Close()

	data, err := json.Marshal(ans)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}

	_, err = f1.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		return
	}
}
