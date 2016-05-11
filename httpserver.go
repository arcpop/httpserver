package main

import (
    "net"
    "net/http"
    "fmt"
	"time"
	"log"
	"strconv"
)

var response = "<html><head><title>Test</title></head><body><p><h2>Page</h2></p></body></html>"

type CustomHandler struct {
    
}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    w.Header().Add("Content-Length", strconv.Itoa(len(response)))
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, response)
}


func main() {
    
    ln, err := net.Listen("tcp4", ":http")
    if err != nil {
        log.Fatal(err)
    }
    s := &http.Server{
        Addr: ":http",
        Handler: &CustomHandler{},
        ReadTimeout: 60 * time.Second,
        WriteTimeout: 60 * time.Second,
        MaxHeaderBytes: 2048,
    }
    s.SetKeepAlivesEnabled(true)
    log.Fatal(s.Serve(ln))
}
