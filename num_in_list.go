package main

// NumInList returns true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if len(list) <= 0 {
		return false
	}
	// initially had this loop without the index
	// placeholder and realised I was looping over
	// indexes instead of values.
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
