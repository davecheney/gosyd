package main // OMIT

import ( // OMIT
	"bufio" // OMIT
	"bytes" // OMIT
	"fmt"   // OMIT
	"os"    // OMIT
) // OMIT

func main() {
	var b bytes.Buffer
	b.Write([]byte("Go 1.6 is awesome\n"))

	// MaxScanTokenSize is 64 * 1024
	// before 1.6 was a hard limit
	longLine := make([]byte, bufio.MaxScanTokenSize+1)
	b.Write(longLine)

	scanner := bufio.NewScanner(&b)
	// Without increasing the buffer, scanner.Err() returns
	// the error "reading long lines bufio.Scanner: token too long"
	// scanner.Buffer allow to increase the limit
	scanner.Buffer(make([]byte, 100), 2*bufio.MaxScanTokenSize)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading long lines", err)
	}
}
