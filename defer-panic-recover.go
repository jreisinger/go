// Defer-panic-recover shows how to recover from a run time panic.
package main

import (
    "log"
    "math/rand"
    "time"
)

func main() {
    protect(mightPanic)
}

// MightPanic sometimes panics, i.e. it crashes the program.
func mightPanic() {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    n := r.Intn(2)
    if n == 0 {
        panic("Panic!")
    }
}

// Protect protects the program from crashing.
func protect(f func()) {
    defer func() {
        log.Println("done") // this will run always
        if x := recover(); x != nil {
            log.Printf("run time panic: %v", x)
        }
    }()
    log.Println("start")
    f()
}
