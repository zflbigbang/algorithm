package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	floatValue := 3.14159265359
	fmt.Fprintf(w, "%.6f\n", floatValue)
}
