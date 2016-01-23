package main

import (
	"fmt"
	"time"
)

func main() {
	_, err := time.Parse(time.RFC822, "29 Feb 15 15:04 MST")
	fmt.Println(err)
}
