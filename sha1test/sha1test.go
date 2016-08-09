package main

import (
    "crypto/sha1"
	"time"
	"fmt"
	"sync/atomic"
)



func main() {
    var block[64]byte
    var counter uint64
    ticker := time.NewTicker(time.Second)
    for i := 0; i < 3; i++ {
        go func() {
            h := sha1.New()
            for {
                h.Sum(block[:])
                atomic.AddUint64(&counter, 1)
            }
        }()
    }
    for {
        h := sha1.New()
        select {
            case <- ticker.C:
                fmt.Println("Hashes calculated: ", atomic.SwapUint64(&counter, 0))  
            default:
                h.Sum(block[:])
                atomic.AddUint64(&counter, 1)
        }
    }
}