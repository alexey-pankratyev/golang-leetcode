/*
Copyright © 2022 aleksei alexey.pankratev@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// singlenumberCmd represents the singlenumberCmd command
var (
	rangeNums []int
)

var	singlenumberCmd = &cobra.Command{
			  Use:   "singlenumber  -r 1,1,3,3,4,5,5 ",
			  Short: "Implementation of the problem of Single Number for litcode",
			  Long: `https://leetcode.com/problems/single-number/
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.	
You must implement a solution with a linear runtime complexity and use only constant extra space.`,
	Run: func(cmd *cobra.Command, nums []string) {
		if verbose {
			fmt.Printf("Verbose mode: nums %v \n", rangeNums)
		}
		singleNumber(rangeNums)
	},
}

func singleNumber(nums []int) int {
	var numberResult int 
	var slice []int
	var found bool
	for i := 0; i < len(nums); i++ {
		slice = append(slice,nums[:i]...)
		slice = append(slice,nums[i+1:]...)
		found := false		
		if verbose {
			fmt.Printf("Index: %v -- Cut a slice: %v\n", i, slice)
		}
		for _, x := range slice {
			if nums[i] == x {	
				found = true
				break	
			}
		}
		if !found {
			numberResult = nums[i] 
		}
		slice = nil
	}
	if !found {
		fmt.Printf("Element is present in the array it's %v \n", numberResult)
	}
    return numberResult
}

func init() {
	rootCmd.AddCommand(singlenumberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// singlenumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
    singlenumberCmd.Flags().IntSliceVarP(
        &rangeNums, "nums", "n", []int{2,2,4,4,1,3,3}, 
        "Numbers",
    )
	singlenumberCmd.MarkFlagRequired("nums")	
}