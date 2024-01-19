package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
    operations := []string{"+", "-", "*", "/"}
    fmt.Printf("Enter a number between 1 and 5: ")
    var level string
    _, err := fmt.Scan(&level)
    if err != nil {
        return
    }
    if len(level) > 1 {
        fmt.Println("You should only enter one digit.")
        return
    }
    numLevel, err := strconv.Atoi(level)
    if err != nil {
        return
    }
    if numLevel > 5 || numLevel < 1 {
        fmt.Println("The level should be between 1 and 5.")
        return
    }
    fmt.Printf("You entered: %s\n", level)

    fmt.Printf("Enter an operation (+, -, *, /): ")
    _, err = fmt.Scan(&level)
    if err != nil {
        return
    }
    if len(level) > 1 {
        fmt.Println("You should only enter one digit.")
        return
    }

    if !slices.Contains(operations, level) {
        fmt.Println("You entered an invalid operation")
        return
    }
        
}

// Generate random numbers according to level
// First number should be larger than second
// Level 1: num1: 1 digit, num2: 1 digit
// Level 2: num1: 2 digit, num2: 1 digit
// Level 3: num1: 2 digit, num2: 2 digit
// Level 4: num1: 3 digit, num2: 1 digit
// Level 5: num1: 3 digit, num2: 2 digit
