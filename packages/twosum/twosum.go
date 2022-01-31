package twosum

import "fmt"

func TwoSum(nums []int, target int) []int {
    mymap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        fmt.Printf("indx:%v target %v, nums: %v, mymap: %v",i,target,nums[i], mymap)
        j, ok := mymap[target-nums[i]]
        fmt.Printf("%d: %v \n",j,ok )
        if ok {
            result := []int{j, i}
            return result
        }
        mymap[nums[i]] = i
    }
    result := []int{-1, -1}
    return result
}