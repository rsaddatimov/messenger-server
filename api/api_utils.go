package api

import (
    "fmt"
    "net/http"
)

func wrongMethod(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Request has wrong method. Use POST request instead")
    
    w.WriteHeader(http.StatusBadRequest)
}

func badRequest(w http.ResponseWriter, r *http.Request, err error) {
    fmt.Fprintln(w, "Error while parsing json")
    fmt.Fprintln(w, err)

    w.WriteHeader(http.StatusBadRequest)
}