package greet

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SayHello() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name:  ")
	scanner.Scan()
	name := strings.Title(scanner.Text())
	fmt.Printf("Hello, %v!\n", name)
}
