package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "testing"
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
    fname := filepath.Join(os.TempDir(), "stdout")
    old := os.Stdout // keep backup of the real stdout
    temp, _ := os.Create(fname) // create temp file
    os.Stdout = temp

    persona := NewPerson(10)
    persona.amIOld()

    temp.Close()
    os.Stdout = old

    out, _ := ioutil.ReadFile(fname)
    out_string := string(out)
    fmt.Println(out_string)
    if strings.TrimRight(out_string, "\n") != "You are young." {
        t.Errorf("MessageAge has a wrong message. Message: %s", out_string)
    }
}