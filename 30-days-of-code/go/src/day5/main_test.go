package main

import (
    "testing"

    "capturer"
)

func TestPrintTable (t *testing.T) {
    valueToPrint := 2
    expectedResult := (
        "2 x 1 = 2\n" +
        "2 x 2 = 4\n" +
        "2 x 3 = 6\n" +
        "2 x 4 = 8\n" +
        "2 x 5 = 10\n" +
        "2 x 6 = 12\n" +
        "2 x 7 = 14\n" +
        "2 x 8 = 16\n" +
        "2 x 9 = 18\n" +
        "2 x 10 = 20\n")

    stdoutCapturer := capturer.NewStdoutCapturer()
    stdoutCapturer.StartCapture()

    printTable(valueToPrint)

    stdoutCapturer.StopCapture()
    printed := stdoutCapturer.GetCapturedString()

    if printed != expectedResult {
        t.Errorf(
            "Printed table is not the " +
            "same as expected:\n%v", printed,
        )
    }
}
