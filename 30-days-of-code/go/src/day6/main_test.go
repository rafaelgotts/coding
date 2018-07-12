package main

import "testing"

func TestGetEvenChars (t *testing.T) {
    input := "Cowboys from Hell"
    expectedChars := "Cwosfo el"

    response := getEvenChars(input)

    if response != expectedChars {
        t.Errorf(
            "EvenChars received differ from expected.\n" +
            "Expected: %s\n" +
            "Received: %s",
            expectedChars,
            response,
        )
    }
}

func TestGetOddChars (t *testing.T) {
    input := "Cowboys from Hell"
    expectedChars := "oby rmHl"

    response := getOddChars(input)

    if response != expectedChars {
        t.Errorf(
            "EvenChars received differ from expected.\n" +
            "Expected: %s\n" +
            "Received: %s",
            expectedChars,
            response,
        )
    }
}
