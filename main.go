package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
)
func main() {
    port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
    listen := net.JoinHostPort("", port)
    log.Println("Start listening on", listen)
    http.HandleFunc("/", HelloServer)
    log.Fatal(http.ListenAndServe(listen, nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "CI/CD Demo! %s", r.URL.Path[1:])
}
