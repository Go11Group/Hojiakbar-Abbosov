package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	client := http.Client{}
	Get(&client)
	Post(&client)
	Put(&client)
	Delete(&client)

}


func Get(client *http.Client) {
	req, err := http.NewRequest("GET", "http://localhost:8080/users/", nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
}
func Post(client *http.Client) {
	body := `
	{
		"name": "Alice Johnson",
		"email": "alice.johnson@example.com",
		"birthday": "2000-12-12 00:00:00",
		"password": "password789"
	}`
	preq, err := http.NewRequest("POST", "http://localhost:8080/users/", bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}
	preq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(preq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	e, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(e))
}
func Put(client *http.Client) {
	body := `
	{
		"name": "Alice Johnson",
		"email": "alice.johnson@example.com",
		"birthday": "2000-12-12 00:00:00",
		"password": "password789"
  	}`
	ureq, err := http.NewRequest("PUT", "http://localhost:8080/users/70efb5e3-58b9-410e-93b7-2b481dab0d3e/update", bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}
	ureq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(ureq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	e, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(e))
}
func Delete(client *http.Client) {
	dreq, err := http.NewRequest("DELETE", "http://localhost:8080/users/70efb5e3-58b9-410e-93b7-2b481dab0d3e/delete", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(dreq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	e, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(e))
}