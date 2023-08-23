package main

import (
	"sort"
)

func main() {

}
func canConstruct(ransomNote string, magazine string) bool {
	hash := [26]int{}
	for i := 0; i < len(magazine); i++ {
		hash[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		hash[ransomNote[i]-'a']--
		if hash[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var rs [][]int
	for k, num := range nums {
		//剪枝
		if num > target && num >= 0 {
			break
		}
		//去重
		if k > 0 && num == nums[k-1] {
			continue
		}
		for i := k + 1; i < len(nums)-2; i++ {
			//二次剪枝
			if nums[i]+num > target && nums[i]+num >= 0 {
				break
			}
			//二次去重
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			//双指针
			l := i + 1
			r := len(nums) - 1
			for l < r {
				sum := num + nums[i] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					rs = append(rs, []int{num, nums[i], nums[l], nums[r]})
					//去重
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return rs
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var rs [][]int
	for i, num := range nums {
		if num > 0 && i == 0 {
			return rs
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := num + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				rs = append(rs, []int{num, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return rs
}
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := map[int]int{}
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			tmp := -(v3 + v4)
			if index, ok := m[tmp]; ok {
				count += index
			}
		}
	}
	return count
}
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		tmp := target - num
		if index, ok := m[tmp]; ok {
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}
func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]struct{}{}
	res := []int{}
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func commonChars(words []string) []string {
	hash := [26]int{}
	res := make([]string, 0)
	for i := 0; i < len(words[0]); i++ {
		hash[words[0][i]-'a']++
	}
	for i := 1; i < len(words); i++ {
		hashTmp := [26]int{}
		for j := 0; j < len(words[i]); j++ {
			hashTmp[words[i][j]-'a']++
		}
		for j := 0; j < 26; j++ {
			a := hash[j]
			b := hashTmp[j]
			hash[j] = min(a, b)
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < hash[i]; j++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func isAnagram(s string, t string) bool {
	hash := [26]int{}
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		hash[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
