//dup reads from the stdin and return the number and content of any duplicated lines

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("line: %s has cardinality %d\n", line, n)
		}
	}
}
