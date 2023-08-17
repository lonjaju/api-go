package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UA: %s\nIP: %s\nX-Forwarded-For: %s\nX-Real-Ip: %s\n",
			r.UserAgent(),
			r.RemoteAddr,
			r.Header.Get("X-Forwarded-For"),
			r.Header.Get("X-Real-Ip"),
		)
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
