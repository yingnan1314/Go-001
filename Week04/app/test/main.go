package test

func main() {
	twoSum([]int{2,5,11,3}, 5)

}

func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i,k := range nums {
		if value,exist := m[target-k];exist {
			result = append(result,value)
			result = append(result,i)
		}
		m[k] = i
	}
	return result
}
