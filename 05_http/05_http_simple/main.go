package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Serve the resource.
			fmt.Fprint(w, "Hello, playground")
		case http.MethodPost:
			// Create a new record.
		case http.MethodPut:
			// Update an existing record.
		case http.MethodDelete:
			// Remove the record.
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Create a new record.
			log.Println("post")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	// go func() {
	log.Fatal(http.Serve(l, nil))
	// }()

	// log.Println("Sending request...")
	// res, err := http.Get("http://localhost:8080/hello")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Reading response...")
	// if _, err := io.Copy(os.Stdout, res.Body); err != nil {
	// 	log.Fatal(err)
	// }
}

// func Main2() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello, playground")
// 	})

// 	log.Println("Starting server...")
// 	go func() {
// 		log.Fatal(http.ListenAndServe("localhost:8080", nil))
// 	}()

// 	log.Println("Sending request...")
// 	res, err := http.Get("http://localhost:8080/hello")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Reading response...")
// 	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
// 		log.Fatal(err)
// 	}
// }
