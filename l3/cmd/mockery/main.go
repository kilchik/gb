package main

import (
)



func main()  {
	
}

func sum(items []uint64) uint64 {
	var total uint64
	for _, i := range items {
		total += i
	}
	return total
}

