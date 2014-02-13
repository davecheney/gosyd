package main 

import "fmt"

func hex(v int) bool {
	return v & 0x8000 == 0x8000
}

func oct(v int) bool {
	return v & 100000 == 100000
}

func main() {
	v := 1 << 15
	fmt.Println(hex(v), oct(v))
}
