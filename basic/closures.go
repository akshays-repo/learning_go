package main

import "fmt"

func withOutIntSeq() int {
	i := 0
	i++
	return i
}

func intSeq() func() {
	i := 0
	return func() {
		i++
	}
}

func main() {

	nextInt := intSeq()

	wnextInt := withOutIntSeq()
	nextInt()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(wnextInt)
	fmt.Println(wnextInt)
	fmt.Println(wnextInt)

	newInts := intSeq()
	fmt.Println(newInts())
}
