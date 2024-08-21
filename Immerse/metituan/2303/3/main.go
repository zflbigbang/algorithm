package main

import (
	"bufio"
	"fmt"

	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, q int
	_, _ = fmt.Fscan(r, &n, &q)
	s := make([]byte, n)
	var rs int
	_, _ = fmt.Fscan(r, &s)
	for i := 0; i < n; i++ {
		if s[i] == 'M' || s[i] == 'T' {
			rs++
		}
	}
	if q < n-rs {
		fmt.Fprint(w, "%d", rs+q)
	} else {
		fmt.Fprint(w, "%d", n)
	}
}
