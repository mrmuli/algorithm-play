package main

// Sum returns the sum of elements in a given slice.
func Sum(sl []int) int {
	count := 0
	for _, i := range sl {
		count += i
	}
	return count
}

// Ongeza returns the sum of elements in a recusive way.
// Fancy way of achieving this:
func Ongeza(sl []int) int {
	// because we need the first index value, confirm it's present.
	if len(sl) == 0 {
		return 0
	}
	// this simply breaks down the slice, tbh makes it more readable
	// with a potential for a stack overflow, careful with it.
	return sl[0] + Ongeza(sl[1:])
}
