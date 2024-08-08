import "fmt"

func simpleVariablesExample() {
    // Declare with var
    var age int = 30
    // Declare with const
    const isMarried bool = true or false
    
    // Declare with shorthand
    name := "Bob"
    isStudent := true

    // Print values
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Is Student:", isStudent)
}
