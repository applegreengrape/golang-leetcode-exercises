// https://leetcode.com/problems/two-sum/

package hashtable

func TwoSum1(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i:= range nums {
        m[i] = nums[i]
    }

	for id1, val := range m {
		check := target - val 
		for id2, val := range m {
			if (val == check) {
				if ( id1 != id2 ){
					return []int{id1, id2}
				}
			}
		}
	}

	return nil 
}

func TwoSum2(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i:= range nums {
        m[i] = nums[i]
    }

	for id1, val1 := range m {
		for id2, val2 := range m {
			if ( val1 + val2 == target){
				if ( id1 != id2 ){
					return []int{id1, id2}
				}
			}
		}
	}

	return nil 
}