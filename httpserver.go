package main

import (
    "net/http"
    "fmt"
	"time"
)

var response = "<html><head><title>Test</title></head><body><p><h2>Page</h2></p></body></html>"

type CustomHandler struct {
    
}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprint(w, response)
}


func main() {
    s := &http.Server{
        Addr: ":http",
        Handler: &CustomHandler{},
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 5 * time.Second,
        MaxHeaderBytes: 2048,
    }
    s.ListenAndServe()
}
