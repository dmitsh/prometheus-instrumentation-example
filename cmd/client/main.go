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
		path := paths[rand.Intn(len(paths))]
		resp, err := http.Get(server + path)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		resp.Body.Close()
		time.Sleep(100 * time.Millisecond)
	}

}
