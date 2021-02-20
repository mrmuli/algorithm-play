package main

func Reverse(nm string) string {
	// Feels over-engineered tbh...
	// 2 slices, one to append string chars
	// a final one for the reversed string.
	var uno []string
	var bruno []string
	for _, j := range nm {
		uno = append(uno, string(j))
	}
	// end defines the last index of sorted slice
	// loops through the slice and reverses into 2nd slice
	end := len(uno) - 1
	for k := 0; k < len(uno); k++ {
		bruno = append(bruno, uno[end])
		end--
	}
	// This final step loops through reversed string
	// and concatenates each char into a string val.
	var val string
	for i := 0; i < len(bruno); i++ {
		val += string(bruno[i])
	}
	return val
}
