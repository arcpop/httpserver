package main

import (
    "net/http"
    "fmt"
	"time"
)

var serverAddresses = []string{"192.168.200.20:8080","192.168.200.30:8080"}

var response = "<html><head><title>Test</title></head><body><p><h2>Page</h2></p></body></html>"

type CustomHandler struct {
    
}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprint(w, response)
}


func main() {
    s1 := &http.Server{
        Addr: serverAddresses[0],
        Handler: &CustomHandler{},
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 5 * time.Second,
        MaxHeaderBytes: 2048,
    }
    go s1.ListenAndServe()
    s2 := &http.Server{
        Addr: serverAddresses[1],
        Handler: &CustomHandler{},
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 5 * time.Second,
        MaxHeaderBytes: 2048,
    }
    s2.ListenAndServe()
}
