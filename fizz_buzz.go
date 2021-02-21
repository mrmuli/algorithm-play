package main

import "fmt"

// LizzBuzz will print out all numbers from 1
// to N but when a number is divisible by 3
// the output is "Fizz" and "Buzz" when divisible by
// 5. For numbers divisible by both, "FizzBuzz".
// Expected is a print of all the values in range. i.e.
// "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz"
// This one doesn't satisfy test case requirements. Check below.
func LizzBuzz(num int) {
	var nums []string
	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 != 0 {
			nums = append(nums, "FizzBuzz")
		} else if i%3 == 0 {
			nums = append(nums, "Fizz")
		} else if i%5 == 0 {
			nums = append(nums, "Buzz")
		} else {
			n := fmt.Sprintf("%v", i)
			nums = append(nums, n)
		}
	}
	var val string
	for i := range nums {
		if len(nums) > 1 {
			val += string(nums[i]) + ","
		} else {
			val += string(nums[i])
		}
	}
	fmt.Print(val)
}

// The challenge in this algorithm is comma and space placement...

// FizzBuzz is a working example of LizzBuzz
func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		HandleN(i)
		fmt.Print(", ")
	}
	HandleN(n)
	fmt.Println()
}

// HandleN handles the final value due to test case requirements.
// Without such requirements, the FizzBuzz loop would evaluate
// the final value i.e. 15 in this case: FizzBuzz(15) so we have
// to separate it's evaluation and print a new line at the end.
func HandleN(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Print("Fizz Buzz")
	} else if n%3 == 0 {
		fmt.Print("Fizz")
	} else if n%5 == 0 {
		fmt.Print("Buzz")
	} else {
		fmt.Print(n)
	}
}
