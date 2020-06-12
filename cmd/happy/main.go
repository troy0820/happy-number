package main

import (
	"fmt"

	"github.com/troy0820/happy-number/pkg/happy"
)

func main() {
	if err := happy.Run(); err != nil {
		fmt.Println(err)
	}
}
