/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// twosumCmd represents the twosum command
var (
	target int
	rangeFlagNumbers []string
)

var	twosumCmd = &cobra.Command{
			  Use:   "twosum 1,2,3,4,5 --target 8",
			  Short: "Implementation of the problem of two-sums for litcode",
			  Long: `https://leetcode.com/problems/two-sum/ 
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.`,
	Run: func(cmd *cobra.Command, nums []string) {
		var twosumArr []int
		for _, ival := range rangeFlagNumbers {
			itemp, _ := strconv.Atoi(ival)
			twosumArr = append(twosumArr,itemp)
		  }
		result := TwoSum(twosumArr,target)
		fmt.Printf("result: %v \n", result)
	},
}

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
    twosumCmd.Flags().StringSliceVarP(
        &rangeFlagNumbers, "range", "r", []string{"1:10"}, 
        "Range of numbers. Optional",
    )
	twosumCmd.MarkFlagRequired("rangeFlagNumbers")	
}
