/*
Copyright © 2022 aleksei alexey.pankratev@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twosumCmd represents the twosum command
var (
	target int
	rangeFlagNumbers []int
)

var	twosumCmd = &cobra.Command{
			  Use:   "twosum -r 1,2,3,4,5 -t 8 or twosum -r 1:5 -t 8",
			  Short: "Implementation of the problem of Two Sums for litcode",
			  Long: `https://leetcode.com/problems/two-sum/ 
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.`,
	Run: func(cmd *cobra.Command, nums []string) {
		if verbose {
			fmt.Printf("verbose: rangeFlagNumbers %v \n", rangeFlagNumbers)
		}
		result := TwoSum(rangeFlagNumbers,target)
		fmt.Printf("result: %v \n", result)
	},
}

func TwoSum(nums []int, target int) []int {
    mymap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
		j, ok := mymap[target-nums[i]]	
		if verbose {
			fmt.Printf("arr_mymap=[%d %v] ",j,ok )
        	fmt.Printf("indx:%v target %v, nums: %v, mymap: %v\n",i,target,nums[i], mymap)			
		}       
        if ok {
            result := []int{j, i}
            return result
        }
        mymap[nums[i]] = i
    }
    result := []int{-1, -1}
    return result
}

func init() {
	rootCmd.AddCommand(twosumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// twosumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	twosumCmd.Flags().IntVarP (&target, "target","t", 0, "integer target")
	twosumCmd.MarkFlagRequired("target")
    twosumCmd.Flags().IntSliceVarP(
        &rangeFlagNumbers, "range", "r", []int{1:10}, 
        "Range of numbers. Optional",
    )
	twosumCmd.MarkFlagRequired("rangeFlagNumbers")	
}
