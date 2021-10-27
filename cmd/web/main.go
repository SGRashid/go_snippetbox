package main
 
import (
    "log"
    "net/http"
)

const port = ":4000"
 
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)
 
    log.Printf("Сервер запущен на localhost%v", port)
    err := http.ListenAndServe(port, mux)
    log.Fatal(err)
}