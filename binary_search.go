package main

import (
	"math"
)

func BinarySearch(haystack []int, needle int) (bool, int) {
	lo := 0
	hi := len(haystack)
	for lo < hi {
		m := int(math.Floor(
			float64(lo) + (float64(hi-lo) / 2),
		))
		v := haystack[m]
		if v == needle {
			return true, m
		} else if v < needle {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return false, -1
}
