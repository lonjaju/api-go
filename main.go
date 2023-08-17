package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UA: %s\nIP: %s\nX-Forwarded-For: %s\nX-Real-Ip: %s\n",
			r.UserAgent(),
			r.RemoteAddr,
			r.Header.Get("X-Forwarded-For"),
			r.Header.Get("X-Real-Ip"),
		)
	})
	http.ListenAndServe(":8888", nil)
}
