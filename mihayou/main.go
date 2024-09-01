package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type ntos struct {
	n, s int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var n int
	defer w.Flush()
	fmt.Fscanln(r, &n)
	ntoss := make([]ntos, n)
	var tmp, sum int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &tmp)
		fmt.Fprintf(w, "%d ", tmp)
		sum = getSum(tmp)
		ntoss[i] = ntos{tmp, sum}

	}

	//对ntos 进行快排，根据 s 进行排序
	sort.Slice(ntoss, func(i, j int) bool {
		return ntoss[i].s < ntoss[j].s
	})
	for _, v := range ntoss {
		fmt.Fprintf(w, "%d ", v.n)
	}

}

// 求一个数的各位之和
func getSum(num int) int {
	var sum int
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
