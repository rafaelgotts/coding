package main

import (
    "reflect"
    "testing"
)

func TestPrintReversedArray (t *testing.T) {
    input := []string{"1", "4", "3", "2"}
    expectedOutput := []string{"2", "3", "4", "1"}

    result := reverseArray(input)

    if !reflect.DeepEqual(result, expectedOutput) {
        t.Errorf(
            "Result differs from expectedOutput.\n" +
            "Expected: %s\n" +
            "Result: %s",
            expectedOutput,
            result,
        )
    }
}
