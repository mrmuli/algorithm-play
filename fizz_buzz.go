package main

import "fmt"

// FizzBuzz will print out all numbers from 1
// to N but when a number is divisible by 3
// the output is "Fizz" and "Buzz" when divisible by
// 5. For numbers divisible by both, "FizzBuzz".
// Expected is a print of all the values in range. i.e.
// "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz"
func FizzBuzz(num int) {
	var nums []string
	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			nums = append(nums, "Fizz")
		} else if i%5 == 0 {
			nums = append(nums, "Buzz")
		} else if i%3 == 0 && i%5 != 0 {
			nums = append(nums, "FizzBuzz")
		} else {
			n := fmt.Sprintf("%v", i)
			nums = append(nums, n)
		}
	}
	var val string
	for i := range nums {
		if len(nums) > 1 {
			val += string(nums[i]) + ", "
		} else {
			val += string(nums[i])
		}
	}
	fmt.Println(val)
}

// The challenge in this algorithm is comma and space placement...
