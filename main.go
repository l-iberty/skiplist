package main

import "github.com/liberty/skiplist/skiplist"

func compare(a, b interface{}) int {
	var x int32
	var y int32

	if a != nil {
		x = a.(int32)
	}
	if b != nil {
		y = b.(int32)
	}

	if x > y {
		return +1
	} else if x < y {
		return -1
	}
	return 0
}

func main() {
	list := skiplist.NewSkipList(compare)
	list = list
}
