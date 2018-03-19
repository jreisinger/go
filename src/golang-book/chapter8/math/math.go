package math

func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Find a minimun value in a slice of floats
func Min(xs []float64) float64 {
    var min float64
    min, xs = xs[0], xs[1:]
    for _, x := range xs {
        if x < min {
            min = x
        }
    }
    return min
}

// Find a maximum value in a slice of floats
func Max(xs []float64) float64 {
    var max float64
    for i, x := range xs {
        if i == 0 || x > max {
            max = x
        }
    }
    return max
}
