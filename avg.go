package main

import "fmt"

func main() {
    nums := []float64{1,2,3,4,5}

    sum := 0.0
    for _, v := range nums {
        sum += v
    }
    fmt.Println(sum / float64(len(nums)))
}
