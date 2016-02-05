package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}

	_, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("--->", os.Args[1], req.URL.String())
	})

	http.ListenAndServe(":"+os.Args[1], nil)
}
