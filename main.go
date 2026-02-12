package main
import (
    "fmt"
    "net/http"
    "os"
)
func home(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(w, "Hello from the DevOps Container! Host: %s\n", hostname)
}
func main() {
    http.HandleFunc("/", home)
    fmt.Println("Server starting on :8080...")
    http.ListenAndServe(":8080", nil)
}
