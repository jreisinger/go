package main

import (
	"fmt"
	"poetry_stats"
)

func main() {
    p := poetry_stats.Poem{{"Line1", "Line2", "Line3"}}
    v, c, puncs := p.Stats()
    fmt.Printf("vowels: %d, consonants: %d, puncs: %d\n", v, c, puncs)
    fmt.Printf("stanzas: %d, lines: %d\n", p.NumStanzas(), p.NumLines())
}
