package main

import (
    "testing"

    "capturer"
)

func TestShouldCreatePersonStruct(t *testing.T) {
    _ = NewPerson(0)
}

func TestPersonAgeMustBeSetOnInitialization(t *testing.T) {
    persona := NewPerson(10)

    if persona.age != 10 {
        t.Error("Person age has not set on initialization")
    }
}

func TestNegativeAgeMustSetToZero(t *testing.T) {
    persona := NewPerson(-10)

    if persona.age < 0 {
        t.Error("Negative age was accepted on initialization")
    }
}

func TestYearPassesShouldIncrementOneYear(t *testing.T) {
    persona := NewPerson(10)
    persona.yearPasses()

    if persona.age != 11 {
        t.Errorf("Age is not the same expected after yearPasses call. Age %d", persona.age)
    }
}

func TestAmIOldSholdChangeMessageAge(t *testing.T) {
    stdoutCapturer := capturer.NewStdoutCapturer()
    stdoutCapturer.StartCapture()

    persona := NewPerson(10)
    persona.amIOld()

    stdoutCapturer.StopCapture()

    out_string := stdoutCapturer.GetCapturedString()
    if out_string != "You are young.\n" {
        t.Errorf("MessageAge has a wrong message. Message: %s", out_string)
    }
}