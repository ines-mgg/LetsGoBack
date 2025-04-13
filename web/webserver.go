package main

import (
    "log"
    "net/http"
)

func main() {
    // Sert les fichiers statiques depuis le dossier courant
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    log.Println("Serveur lancé sur http://localhost:5500")
    err := http.ListenAndServe(":5500", nil) // Start the web server
    if err != nil {
        log.Fatal(err)
    }
}