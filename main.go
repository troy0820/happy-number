package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var collector []int

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter number to see if it's Happy: ")
	reader.Scan()
	text := reader.Text()
	text = strings.Replace(text, " ", "", -1)
	happy, _ := strconv.Atoi(text)
	fmt.Printf("Checking to see if %v is Happy \n", happy)
	HappyNumber(happy)
}

func HappyNumber(n int) bool {
	sl := Parse(n)
	var sum int
	var answers []int
	for _, v := range sl {
		sum += v * v
	}
	answers = append(answers, sum)
	answerSum := Parse(sum)
	length := len(answerSum)
	if answerSum[length-1] == 1 {
		fmt.Printf("This is a Happy Number :D \n")
		return true
	}
	num, err := strconv.Atoi(ParseToString(answers))
	if err != nil {
		fmt.Printf("This is an error parsing to string %v \n", err)
		return false
	}
	if checkInSlice(collector, sum) == true {
		fmt.Printf("This is sad Number :(  \n")
		return false
	}
	//	fmt.Println("sum answers: ", answers)

	collector = append(collector, sum)
	HappyNumber(num)

	return true
}
func checkInSlice(sl []int, num int) bool {
	for _, v := range sl {
		if v == num {
			return true
		}
	}
	return false
}
func ParseToString(sl []int) string {
	var thing strings.Builder
	for _, v := range sl {
		str := strconv.Itoa(v)
		thing.WriteString(str)
	}
	return thing.String()
}
func Parse(n int) []int {
	num := strconv.Itoa(n)
	var newNums []int
	for _, v := range num {
		number, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("You had an error")
		}
		newNums = append(newNums, number)
	}
	//	fmt.Println("This is the new array", newNums)
	return newNums
}

//pseudo cod

/*
  check if it is 1 or ends in 1 and return true it's happy
  parse input to find length, put in list, square list, and then sum <-- own function
  check if number returned from function is equal to 1, if not put new number into function
  store answer in list to check if the number is repeated to see if you are in a loop.



*/
