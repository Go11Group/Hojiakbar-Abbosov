package main

import "fmt"

func main() {

	m := make(map[string]map[string]int)

	m["book"] = map[string]int{}
	m["book"]["Tohir Malik"] = 10

	m["copybook"] = map[string]int{}
	m["copybook"]["Alvido bolalik"] = 320

	for key, val := range m  {
		fmt.Println(key,val)
	}

}