package main
import "fmt"

func removeElement(nums []int, val int) int {
    numOfElementsNotEqual:= 0;
    for i,_ := range nums {
		if nums[i] != val{
			nums[numOfElementsNotEqual] = nums[i]
			numOfElementsNotEqual++
			
		}
	}
	fmt.Printf("Array: %+v", nums)
	return numOfElementsNotEqual
}

func removeDuplicateSortedII(nums []int) int {
	isSame := false;
	numOfElementsNotEqual := 0;
	for i,_ := range nums {
		if !isSame {
			nums[numOfElementsNotEqual] = nums[i]
			numOfElementsNotEqual++
			isSame = false
			if i+1 < len(nums) {
				if nums[i] == nums[i+1]{
					nums[numOfElementsNotEqual] = nums[i]
					numOfElementsNotEqual++
					isSame = true
				}
			}
		}else if i+1 < len(nums) {
			if nums[i] != nums[i+1]{
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
	
	for i:= 0; i<len(nums)+1; i++{
		curNum := prevNum
		newLoc := (prevIndex+k) % (len(nums)-1)
		if (prevIndex+k) == (len(nums)-1) {
			newLoc = len(nums)-1
		}else if (prevIndex+k) > (len(nums)-1){
			newLoc--
		}
		prevIndex = newLoc
		prevNum = nums[prevIndex]
		nums[newLoc] = curNum
	}

	fmt.Printf("Array: %+v", nums)
}

func maxProfit(prices []int) int{
	max := 0
	sellDay := len(prices)-1
	buyDay := 0
	for i:= len(prices)-1; i>-1; i--{
		if prices[i] >= prices[sellDay] && i > buyDay {
			sellDay = i
		}
		curmax := prices[sellDay] - prices[i]
		if curmax > max{
			max = curmax
		}
	}

	if max > 0 {
		return max
	}
	return 0	
}

// func maxProfitHard(prices []int) int{
// 	 (go greeedy gang)
// }
func longestPalindromicSubstring(s string) string{
	palindromes := make(map[string]bool)
	longestWord := ""
	for i:=0; i < len(s); i++ {
		if s[i] == s[i+1]{
			longestWord = s[i:i+2]
			palindromes[s[i:i+2]] = true
		}
	}
	cut := 3
	for cut < len(s)+1{
		for i:=0;i < len(s)-cut+1;i++{
			curword := s[i:i+cut]
			if curword[0] == curword[len(curword)-1]{
				if cut != 3 {
					palindromes[curword] = true
					longestWord = curword
				}else {
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
	test_arr :=  []int{3,2,6,5,0,3};

	fmt.Printf("Max Profit: %d", maxProfit(test_arr))
}