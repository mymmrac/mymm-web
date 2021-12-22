package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n\n%d\n\nPath: %q\n", rand.Int(), r.URL)
	})

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

