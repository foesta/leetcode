package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	numOfElementsNotEqual := 0
	for i, _ := range nums {
		if nums[i] != val {
			nums[numOfElementsNotEqual] = nums[i]
			numOfElementsNotEqual++

		}
	}
	fmt.Printf("Array: %+v", nums)
	return numOfElementsNotEqual
}

func removeDuplicateSortedII(nums []int) int {
	isSame := false
	numOfElementsNotEqual := 0
	for i, _ := range nums {
		if !isSame {
			nums[numOfElementsNotEqual] = nums[i]
			numOfElementsNotEqual++
			isSame = false
			if i+1 < len(nums) {
				if nums[i] == nums[i+1] {
					nums[numOfElementsNotEqual] = nums[i]
					numOfElementsNotEqual++
					isSame = true
				}
			}
		} else if i+1 < len(nums) {
			if nums[i] != nums[i+1] {
				isSame = false
			}
		}
	}
	fmt.Printf("Array: %+v", nums)
	return numOfElementsNotEqual
}

func rotateArray(nums []int, k int) {
	prevIndex := 0
	prevNum := nums[prevIndex]

	for i := 0; i < len(nums)+1; i++ {
		curNum := prevNum
		newLoc := (prevIndex + k) % (len(nums) - 1)
		if (prevIndex + k) == (len(nums) - 1) {
			newLoc = len(nums) - 1
		} else if (prevIndex + k) > (len(nums) - 1) {
			newLoc--
		}
		prevIndex = newLoc
		prevNum = nums[prevIndex]
		nums[newLoc] = curNum
	}

	fmt.Printf("Array: %+v", nums)
}

func maxProfit(prices []int) int {
	max := 0
	sellDay := len(prices) - 1
	buyDay := 0
	for i := len(prices) - 1; i > -1; i-- {
		if prices[i] >= prices[sellDay] && i > buyDay {
			sellDay = i
		}
		curmax := prices[sellDay] - prices[i]
		if curmax > max {
			max = curmax
		}
	}

	if max > 0 {
		return max
	}
	return 0
}

func canJump(nums []int) bool {
	curIndex := 0
	length := len(nums)
	prevIndex := curIndex
	if curIndex == length-1 {
		return true
	} else if nums[curIndex] == 0 {
		return false
	} else {
		for curIndex := 0; curIndex <= len(nums); curIndex = curIndex {
			prevIndex = curIndex
			numJumps := nums[curIndex]
			if curIndex+numJumps >= length-1 {
				return true
			}
			for s := 1; s <= numJumps; s++ {
				if nums[curIndex+s] > numJumps-s {
					curIndex = curIndex + s
					break
				}
			}
			if curIndex != prevIndex {
				continue
			}
			return false
		}
	}
	return false
}

func getArrProduct(nums []int) int{
	ans := nums[0]
	for i,v := range(nums){
		if i == 0{
			continue
		}
		ans = ans*v
	}
	return ans
}

func productExceptSelf(nums []int) []int{
	retArr := make([]int, len(nums))
	for i,_ := range(nums){
		leftArr := nums[0:i]
		rightArr := nums[i+1:]
		if len(leftArr) < 1 {
			retArr[i] = getArrProduct(rightArr)
		}else if len(rightArr) < 1 {
			retArr[i] = getArrProduct(leftArr)
		}else {
			retArr[i] = getArrProduct(leftArr) * getArrProduct(rightArr)
		}
	}
	return retArr
}

func prodder(nums []int) []int{
	
	total := nums[0]
	retArrLeft := make([]int, len(nums))
	retArrRight := make([]int, len(nums))
	ansArr := make([]int, len(nums))
	for i,_ := range(nums){
		if i == 0{
			continue
		}
		retArrLeft[i] = total
		total = total*nums[i]
	}

	total = nums[len(nums)-1]
	for i:=len(nums)-1;i>-1;i--{
		if i == len(nums)-1{
			continue
		}
		retArrRight[i]=total
		total = total*nums[i]
	}

	for i,_ := range(nums){
		ansArr[i] = retArrLeft[i] * retArrRight[i]
		if i == 0{
			ansArr[i] = retArrRight[i]
		}
		if i == len(nums)-1{
			ansArr[i] = retArrLeft[i]
		}
	}
	return ansArr

}

func hIndex(citations []int) int {
	length := len(citations)
	buckets := make([]int, length+1)
	for i := 0; i < length; i++ {
		if citations[i] >= length {
			buckets[length]++
		} else {
			buckets[i]++
		}
	}
	count := 0
	for i := length; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}
	return 0
}

//	func maxProfitHard(prices []int) int{
//		 (go greeedy gang)
//	}
func longestPalindromicSubstring(s string) string {
	palindromes := make(map[string]bool)
	longestWord := ""
	for i := 0; i < len(s); i++ {
		if s[i] == s[i+1] {
			longestWord = s[i : i+2]
			palindromes[s[i:i+2]] = true
		}
	}
	cut := 3
	for cut < len(s)+1 {
		for i := 0; i < len(s)-cut+1; i++ {
			curword := s[i : i+cut]
			if curword[0] == curword[len(curword)-1] {
				if cut != 3 {
					palindromes[curword] = true
					longestWord = curword
				} else {
					palindromes[curword] = true
					longestWord = curword
				}
			}
		}
		cut++
	}
	return longestWord

}

func main() {
	test_arr := []int{-1,1,0,-3,3}

	fmt.Printf("Ans: %+v", prodder(test_arr))
}
