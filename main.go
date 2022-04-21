package main

import (
        "encoding/json"
        "github.com/gorilla/mux"
        "log"
        "net/http"
)

func main() {
        handler := mux.NewRouter()
        handler.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode("Pong") })
        server := &http.Server{Addr: "0.0.0.0:8001", Handler: handler}
        log.Println("Starting the server!")
        if err := server.ListenAndServe(); err != nil {
                log.Fatalln("[ERROR] Error while starting the server")
        }
}
