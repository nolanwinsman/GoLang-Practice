package main

import (
	"fmt"
	"github.com/inancgumus/screen"
)

func main() {
	for i := 0; i < 10; i++ {
		screen.Clear()
		fmt.Println(i)
	}
}
