package main

import (
    "log"
    "math/rand"
    "time"
)

func main() {
    protect(mightPanic)
}

func mightPanic() {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    n := r.Intn(2)
    if n == 0 {
        panic("Panic!")
    }
}

func protect(f func()) {
    defer func() {
        log.Println("done")
        if x := recover(); x != nil {
            log.Printf("run time panic: %v", x)
        }
    }()
    log.Println("start")
    f()
}
