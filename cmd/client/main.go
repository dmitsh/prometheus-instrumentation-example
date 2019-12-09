package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	server := "http://localhost:8080"
	paths := []string{"/api", "/home", "/store", "/user"}

	for {
		url := server + paths[rand.Intn(len(paths))]
		fmt.Printf("GET %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println("Response status:", resp.Status)
		resp.Body.Close()
		time.Sleep(100 * time.Millisecond)
	}
}
