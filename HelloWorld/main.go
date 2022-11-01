package main

import "fmt"

// Solving the LeetCode 75 challange on leetcode
// Solutions to all problems are listed below as functions
// Structures
 type ListNode struct {
	Val int
	Next *ListNode
}


// Functions
func runningSum(nums []int) {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

}
func pivotIndex(nums []int) int {

	leftsum := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if leftsum == sum-leftsum-nums[i] {
			return i
		}
		leftsum += nums[i]
	}
	return -1
}
func isIsomorphic(s string, t string) bool {

	Hash1 := make(map[byte]byte)
	Hash2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		val, exist := Hash1[s[i]]
		if !exist {
			Hash1[s[i]] = t[i]
			
		} else {
			if val != t[i] {
				return false
			}
		}

		val1, exist1 := Hash2[t[i]]
		if !exist1 {
			Hash2[t[i]] = s[i]

		} else {
			if val1 != s[i] {
				return false
			}
		}

	}

	return true
}
func isSubsequence(s string, t string) bool {

	done := len(s)
	if done == 0 {
		return true
	}

	for i, k := 0, 0; i < len(t); i++ {
		if t[i] == s[k] {
			done--
			k++
		}
		if done == 0 {
			return true
		}
	}
	return false
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}
	if list1.Val <= list2.Val{
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}else{
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}

func main() {
	s := "axc"
	t := "ahbgdc"

	if isSubsequence(s, t) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
