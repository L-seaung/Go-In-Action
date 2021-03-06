package main

import "fmt"
import "time"

func main(){
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("tow")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    wahtAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println(("I am a bool"))
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't konw type %T\n", t)
        }
    }

    wahtAmI(true)
    wahtAmI(1)
    wahtAmI("hey")
}
