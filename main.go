package main

import (
	"fmt"
	"net/http"
)

func main() {
	portNum := 8003
	addr := fmt.Sprintf(":%d", portNum)

	fmt.Printf("start listening server at %d\n", portNum)

	http.HandleFunc("/", helloWorld)

	if err := http.ListenAndServe(addr, nil); err != nil {
		println("error!!!!!!!!!!")
		panic(err)
	}

	fmt.Println("closing server")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world!")
}
