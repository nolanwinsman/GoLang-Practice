// Go module to practice writing structs

package main

import "fmt"

type Animal struct {  
    name string
	age int
}

func main() {
	tiger := Animal{name: "Tony", age: 40} // creates variable tiger of type Animal
	fmt.Printf("The Tiger's name is %s he is %d years old\n",tiger.name, tiger.age)
}