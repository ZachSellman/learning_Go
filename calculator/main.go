package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")
	//Here I ask for Value1 from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Value1: ")
	input1, _ := reader.ReadString('\n')
	int1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		panic("Error: user did not input a valid number")
	}

	fmt.Printf("Value2: ")
	input2, _ := reader.ReadString('\n')
	int2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		panic("Error: user did not input a valid number")
	}

	sumValue := (int1 + int2)
	sumValue = math.Round(sumValue*100) / 100
	fmt.Printf("The rounded sum of %v and %v is %v\n\n", int1, int2, sumValue)


	
	
	//Here I ask for Value2 from the user

	//Here I use the sum function to add both together, and the convert string function to convert to int
}
