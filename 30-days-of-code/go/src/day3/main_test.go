package main

import "testing"

func TestIsOdd(t *testing.T) {
    t.Run("3 is a odd number", func(t *testing.T) {
        got := isOdd(3)
        want := true

        if got != want {
            t.Errorf("3 should return even, but returned %v", got)
        }
    })
}

func TestIsWeird(t *testing.T) {
    t.Run("3 is odd, then is a weird number!", func(t *testing.T) {
        got := isWeird(3)
        want := true

        if got != want {
            t.Errorf("3 should return true, but returned %v", got)
        }
    })
}
