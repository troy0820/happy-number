package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var collector []int
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter number to see if it's Happy: ")
	reader.Scan()
	//	text := reader.Text()
	//	text = strings.Replace(text, " ", "", -1)
	//	happy, _ := strconv.Atoi(text)
	var happy int
	fmt.Sscanf(reader.Text(), "%d", &happy)
	fmt.Printf("Checking to see if %v is Happy \n", happy)
	HappyNumber(happy)
}

//HappyNumber checks to see if a number is Happy or sad
func HappyNumber(n int) bool {
	sl := Parse(n)
	if sl[len(sl)-1] == 1 {
		fmt.Println(green("This is a Happy Number :D"))
		return true
	}
	var sum int
	var answers []int
	for _, v := range sl {
		sum += v * v
	}
	answers = append(answers, sum)
	answerSum := Parse(sum)
	length := len(answerSum)
	if answerSum[length-1] == 1 {
		fmt.Println(green("This is a Happy Number :D"))
		return true
	}
	num, err := strconv.Atoi(ParseToString(answers))
	if err != nil {
		fmt.Printf("This is an error parsing to string %v \n", err)
		return false
	}
	if checkInSlice(collector, sum) {
		fmt.Println(red("This is sad Number :("))
		return false
	}

	collector = append(collector, sum)
	return HappyNumber(num)
}
func checkInSlice(sl []int, num int) bool {
	for _, v := range sl {
		if v == num {
			return true
		}
	}
	return false
}

//ParseToString takes a slice of ints and returns a string
//ex. [1,2,3] -> "123"
func ParseToString(sl []int) string {
	var thing strings.Builder
	for _, v := range sl {
		str := strconv.Itoa(v)
		thing.WriteString(str)
	}
	return thing.String()
}

//Parse takes an int and makes it a slice of ints of the same number
//ex 123 -> [1,2,3]
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
	return newNums
}

//pseudo code

/*
  check if it is 1 or ends in 1 and return true it's happy
  parse input to find length, put in list, square list, and then sum <-- own function
  check if number returned from function is equal to 1, if not put new number into function
  store answer in list to check if the number is repeated to see if you are in a loop.
*/
