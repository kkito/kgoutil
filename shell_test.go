package kgoutil

import (
    "testing"
)

func TestBash(t *testing.T) {
    result := Bash("ls -l")
    if len(result) < 10 {
        t.Fatal("error return")
    }
}

