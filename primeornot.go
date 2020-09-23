package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	var givenNumber int

	fmt.Print("Enter a positive integer number > ")
	fmt.Scanf("%d", &givenNumber)

	flag, divisors, err := isPrime(givenNumber)

	if err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
		main()
	} else {
		if flag == true {
			fmt.Printf("%d is a prime number.\n", givenNumber)
		} else {
			fmt.Printf("%d is not a prime number.\n", givenNumber)
			fmt.Printf("Divisors for %d: %s\n", givenNumber, divisors)
		}
	}
}

func isPrime(num int) (bool, string, error) {
	var divisors []int
	var divisorsText string
	divisors = append(divisors, 1) // 1 is definitely a divisor for every number
	divisorsText += "1"

	if num < 1 {
		return false, divisorsText, errors.New("not a positive number")
	} else if num == 1 {
		return false, divisorsText, nil
	}
	// As these if and else if blocks end with return, there isn't an else block
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			divisorsText += ", " + strconv.Itoa(i)
		}
	}
	if divisors[1] == num {
		return true, divisorsText, nil // Prime
	}
	return false, divisorsText, nil // Not prime

}
