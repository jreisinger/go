package main

import "fmt"

func main() {
	nums := []float64{1, 2, 3, 4, 5}
	fmt.Println(avg(nums))
}

func avg(nums []float64) float64 {
	//panic("not implemented")
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}
