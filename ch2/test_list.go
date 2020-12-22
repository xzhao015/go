package ch2

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func sum(nums []int) int {
	var sumOuput int
	for _, value := range nums {
		sumOuput += value
	}
	return sumOuput
}

func generateList(nums []int, numList []int) []int {
	fmt.Println(nums)
	fmt.Println(numList)
	if len(nums) == 1 {
		return numList
	}
	numsM := nums[1:]
	var sumList []int
	for i := 0; i < len(nums); i++ {
		sumList = append(sumList, sum(nums[0:i+1]))
	}
	for _, value := range numList {
		sumList = append(sumList, value)
	}
	return generateList(numsM, sumList)

}

//RangeSum is a function.
func RangeSum(nums []int, n int, left int, right int) int {
	numList := generateList(nums, []int{})
	sort.Ints(numList)
	fmt.Println("generate list is ", numList)
	return sum(numList[left-1 : right])
}

//MinDifference is a function.
func MinDifference(nums []int) int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(nums))
		changeto := rand.Int()
		nums[index] = changeto
		fmt.Println(nums)
	}
	sort.Ints(nums)
	return nums[1] - nums[0]

}
