// main.go
package main
import (
    "fmt"
    "net/http"
    "os"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Running on Floci EKS! Version: %s\n", os.Getenv("VERSION"))
    })
    http.ListenAndServe(":8080", nil)
}
