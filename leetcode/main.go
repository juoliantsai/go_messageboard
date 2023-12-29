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

    nums := []int{0,0,1,1,1,2,2,3,3,4}
    res := removeDuplicates(nums)

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
    res := make([]int)
    rr:=0
    for i,v1 := range nums {
        if len(res) == 0 {
            res = append(res,v1)
	} else {
            for j,v2 := range res {
                if v1 == v2 {
                } else {
                }
	    }
	}
    }
    return len(res)
}
