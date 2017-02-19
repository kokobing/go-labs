package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
	}

	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()

	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
