package poetry_stats

import (
    "testing"
)

func TestNumStanzas(t *testing.T) {
    p := Poem{}
    if p.NumStanzas() != 0 {
        t.Fatalf("Empty poem is not empty")
    }

    p = Poem{{"Line1", "Line2", "Line3"}}
    if p.NumStanzas() != 1 {
        t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
    }
    if p.NumLines() != 3 {
        t.Fatalf("Unexpected line count %d", p.NumLines())
    }
}

func TestStats(t *testing.T) {
    p := Poem{}
    v, c, puncs := p.Stats()
    if v != 0 || c != 0 || puncs != 0 {
        t.Fatalf("Bad number of vowels or consonants")
    }

    p = Poem{{"Hello, World!"}}
    v, c, puncs = p.Stats()
    if v != 3 || c != 7 || puncs != 3 {
        t.Fatalf("Bad number of vowels (%d) or consonants (%d)", v, c)
    }
}
