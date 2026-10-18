package main

import "fmt"

func main() {
	var n, menit, total int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&menit)
		total += menit
	}
	jam := total / 60
	menit = total % 60
	fmt.Printf("%d jam %d menit\n", jam, menit)
}
