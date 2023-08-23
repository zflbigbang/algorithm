package main

type H struct {
	id int
}

func main() {

}

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	//1.局部反转
	reverseString(b[0:n])
	reverseString(b[n:])
	//2.整体反转
	reverseString(b)
	return string(b)
}
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
func replaceSpace(s string) string {
	b := []byte(s)
	n := len(s)
	cnt := 0
	for i := 0; i < n; i++ {
		if b[i] == ' ' {
			cnt++
		}
	}
	//扩容
	tmp := make([]byte, cnt*2)
	b = append(b, tmp...)
	j := len(b)
	for i := n; i >= 0; i-- {
		if b[i] != ' ' {
			b[j] = b[i]
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			j -= 3
		}
	}
	return string(b)
}
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	next := make([]int, n)
	getNextCopy(s, next)
	if next[n-1] > 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}
func getNextCopy(s string, next []int) {
	j := 0
	next[j] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return -1
	}
	next := make([]int, n)
	getNext(needle, next)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
func getNext(s string, next []int) {
	j := 0
	next[j] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func reverseWords(s string) string {
	t := []byte(s)
	//1.去除多余空格
	slow := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if slow != 0 {
				t[slow] = ' '
				slow++
			}
			for i < len(s) && s[i] != ' ' {
				t[slow] = s[i]
				slow++
				i++
			}
		}
	}
	t = t[0:slow]
	//2.反转整个字符串
	reverseString(t)
	//3.反转单词
	start := 0
	for i := 0; i <= len(t); i++ {
		if i == len(t) || t[i] == ' ' {
			reverseString(t[start:i])
			start = i + 1
		}
	}
	return string(t)
}

func reverseStr(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			reverseString(t[i : i+k])
		} else {
			reverseString(t[i:len(s)])
		}
	}
	return string(t)
}
