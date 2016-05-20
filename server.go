package main

import(
	"fmt"
	"net/http"
	"log"
)

func main() {
	port := "8080"
	
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("html"))))
	
	i := 0
	http.HandleFunc("/resource/count", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		switch r.Method {
			case http.MethodGet:
				log.Println("Get resource/count ", i)
				fmt.Fprintf(w, "Count: %d", i)
			case http.MethodPost:
				log.Println("Post resource/count ", i)
				i++
			default:
				log.Println("Illegal Method: " + r.Method)
		}
	})
	
	log.Printf("Server startup on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}