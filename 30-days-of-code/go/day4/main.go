package main
import "fmt"

type person struct {
    age int
}

func NewPerson (initialAge int) *person {
    if initialAge < 0 {
        fmt.Println("Age is not valid, setting age to 0")
        initialAge = 0
    }

    return &person{initialAge}
}

func (p *person) yearPasses () {
    p.age++
}

func (p *person) amIOld () {
    if p.age < 13 {
        fmt.Println("You are young.")
    }

    if p.age >= 13 && p.age < 18 {
        fmt.Println("You are a teenager.")
    }

    if p.age >= 18 {
        fmt.Println("You are old.")
    }
}

func main() {
    var T,age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&age)
        p := NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p.yearPasses()
        }
        p.amIOld()
        fmt.Println()
    }
}