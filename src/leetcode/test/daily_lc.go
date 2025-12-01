package test

import (
	"sort"
	"strings"
)



func CountFairPairs(nums []int, lower int, upper int) int64 {
    fair_pairs := int64(0)
    sort.Ints(nums)
    for i := len(nums) -1; i > 0; i -= 1 {

        //upper bound search
        targetHigh := upper - nums[i]
        right_limit := -1
        left_i, right_i := 0, i
        for left_i < right_i {
            m := left_i + (right_i - left_i) / 2
            if nums[m] <= targetHigh {
                right_limit = m
                left_i = m + 1
            } else {
                right_i = m
            }
        }
        //lower bound search
        targetLow := lower - nums[i]
        left_limit := i
        left_i, right_i = 0, i
        for left_i < right_i {
            m := left_i + (right_i - left_i) / 2
            if nums[m] >= targetLow {
                left_limit = m
                right_i = m
            } else {
                left_i = m + 1
            }
        }
        if right_limit >= left_limit {
            fair_pairs += int64(right_limit - left_limit + 1)
        } 
        
    }
    return fair_pairs
}

func TakeCharacters(s string, k int) int {
    n := len(s)
    total := [3]int{}
    for i := range n {
        total[s[i]-'a']++
    }

    if total[0] < k || total[1] < k || total[2] < k {
        return -1
    }

    maxWindow := 0
    window := [3]int{}
    left := 0

    for right := range n {
        window[s[right]-'a']++

        for left <= right && (total[0]-window[0] < k || total[1]-window[1] < k || total[2]-window[2] < k) {
            window[s[left]-'a']--
            left++
        }

        if maxWindow < right-left+1 {
            maxWindow = right - left + 1
        }
    }

    return n - maxWindow
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func CountAndSay(n int) string { 
    var result strings.Builder
    prev_enconding := "1"
    for i := 1; i < n; i = i + 1 {
        for j := 0; j < len(prev_enconding); {
            repeating_times := rune('1')
            actual_value := prev_enconding[j]
            for k := j + 1; k < len(prev_enconding); k = k + 1 {
                curr_val := prev_enconding[k]
                if curr_val != actual_value {
                    break
                }
                repeating_times = 1 + repeating_times
            }
            result.WriteRune(repeating_times)
            result.WriteByte(actual_value)
			j = j + int(repeating_times - '0')
        }
        prev_enconding = result.String()
        result.Reset()
    }

    return prev_enconding
}
