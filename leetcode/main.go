package main

import (
    "fmt"
    "strconv"
)

func main() {
    /*nums := []int{2,7,11,15}
    target := 13
    res := twoSum(nums,target)*/

    /*x := 121
    res := isPalindrome(x)*/

    /*s := "MCMXCIV"
    res := romanToInt(s)*/

    /*nums := []int{0,0,1,1,1,2,2,3,3,4}
    res := removeDuplicates(nums)*/

    /*nums,val := []int{0,1,2,2,3,0,4,2},2
    res := removeElement(nums,val)*/

    /*haystack,needle := "leetcode","leeto"
    res:=strStr(haystack,needle)*/

    /*nums,target := []int{1,3,5,6},5
    res:=searchInsert(nums,target)*/

    s := "   fly me   to   the moon  "
    res:=lengthOfLastWord(s)

    fmt.Println(res)
}
// Two Sum
func twoSum(nums []int, target int) []int {
    cnt := len(nums)-1
    res := make([]int,2)
    for i := 0; i < cnt; i = i+1 {
        for j := i+1; j < cnt; j = j+1 {
            if nums[i] + nums[j] == target {
                res := []int{i,j}
		return res
	    }
	}
    }
    return res
}
// Palindrome Number
func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
    a := []rune(s)
    for i,j := 0,len(a)-1; i < j; i,j = i+1,j-1 {
        a[i],a[j] = a[j],a[i]
    }
    s = string(a)
    c,_ := strconv.Atoi(s)
    fmt.Println(c)
    if x == c {
        return true
    }
    return false
}
// Roman to Integer
func romanToInt(s string) int {
    obj := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
    a := []rune(s)

    r := make([]int, 2)
    for i,j:=0,len(a); i<j;{
        b := string(a[i])
        x := obj[b]
        y := 0
	i = i+1
	if i < j {
            b = string(a[i])
            y = obj[b]
        }
        if x < y {
            r = append(r, y-x)
            i = i+1
        } else {
            r = append(r, x)
        }
    }
    res := 0
    for _,v := range r {
        res = res+v
    }

    return res
}
// Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
    res := make([]int,0)
    for _,v1 := range nums {
        check := true
        if len(res) > 0 {
            for _,v2 := range res {
                if v1 == v2 {
                    check = false
                }
            }
            if check {
                res = append(res,v1)
            }
        } else {
            res = append(res,v1)
        }
    }

    return len(res)
}
// Remove Element
func removeElement(nums []int, val int) int {
    res := make([]int, 0)
    for _,v := range nums {
        if v != val {
            res = append(res, v)
        }
    }
    return len(res)
}
// Find the Index of the First Occurrence in a String
func strStr(haystack string, needle string) int {
    s1 := []rune(haystack)
    s2 := []rune(needle)
    res:= -1
    loop1:
    for i,j := 0,len(s1)-1; i < j; i = i+1 {
        loop2:
        for k,l := 0,len(s2)-1; k < l; k = k+1 {
            if s1[i] == s2[k] {
                if res < 0 {
                    res = i
                }
                i = i+1
                if i > l {
                    break loop1
                }
            } else {
		res = -1
                break loop2
            }
        }
    }

    return res
}
// Search Insert Position
func searchInsert(nums []int, target int) int {
    i,j:=0,len(nums)-1
    for i=0;i<j;i=i+1 {
        if nums[i] >= target {
            return i
        }
    }

    return i+1
}
// Length of Last Word
func lengthOfLastWord(s string) int {
    str := []rune(s)
    word := make([]rune,0)
    words := make([]string,0)
    space := []rune(" ")

    for i,j:=0,len(str); i<j; i=i+1 {
        if str[i] != space[0] {
            word = append(word, str[i])
        }
        if str[i] == space[0] || i+1 == j {
            w := string(word)
            if len(w) > 0 {
                words = append(words, w)
                word = make([]rune,0)
            }
        }
    }

    l:=len(words)
    len:=len(words[l-1])

    return len
}
