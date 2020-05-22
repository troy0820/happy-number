package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {

	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter number to see if it's Happy: ")
	reader.Scan()
	var happy int
	fmt.Sscanf(reader.Text(), "%d", &happy)
	fmt.Printf("Checking to see if %v is Happy \n", happy)
	if num := HappyNumber(happy); num == true {
		fmt.Println("This is a Happy Number")
		return nil
	}
	return fmt.Errorf("This is not a Happy Number")

}

func HappyNumber(n int) bool {
	m := map[int]bool{}
	for {
		if n%10 == 1 {
			return true
		}

		n = processNumber(n)

		if m[n] == true || n == 0 {
			return false
		}

		m[n] = true
	}
}

func processNumber(n int) int {
	var sum, num int
	for 0 < n {
		num = n % 10
		sum += num * num
		n = n / 10
	}
	return sum
}
