package main

import (
	"log"
	"time"
)

// Solving the LeetCode 75 challange on leetcode
// Solutions to all problems are listed below as functions

// Structures
type ListNode struct {
	Val  int
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
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}
func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head
	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	return prev
}
func middleNode(head *ListNode) *ListNode {
	curr := head
	num := 0
	for curr != nil {
		num++
		curr = curr.Next
	}
	curr = head
	index := num / 2

	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr
}
func detectCycle(head *ListNode) *ListNode {

	var visitedMap = make(map[*ListNode]int)
	repeat := false
	i := 0
	for repeat != true {
		if head == nil || head.Next == nil {
			return nil
		}
		_, ok := visitedMap[head]
		if !ok {
			visitedMap[head] = i
		} else {
			return head
		}
		head = head.Next
	}
	return nil
}
func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	over_profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		profit = prices[i] - min
		if over_profit < profit {
			over_profit = profit
		}
	}
	
	return over_profit
}

func solution() {
	test1 := []int{7, 6, 4, 3, 1}
	test2 := []int{7, 1, 5, 3, 6, 4}
	log.Println(maxProfit(test1))
	log.Println(maxProfit(test2))

}
func main() {

	start_time := time.Now()

	solution()
	end_time := time.Since(start_time)
	log.Println("Program finished execution after: ", end_time.Seconds(), "seconds")

}
