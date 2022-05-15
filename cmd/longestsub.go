/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

var (
	inputstr string
)

// longestsubCmd represents the longestsub command
var longestsubCmd = &cobra.Command{
	Use:   "longestsub -s examplestring",
	Short: "Implementation of the problem of Longest Substring Without Repeating Characters for litcode",
	Long: `https://leetcode.com/problems/longest-substring-without-repeating-characters/ 
Given a string s, find the length of the longest substring without repeating characters.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := LongestSub(inputstr)
		fmt.Printf("%v \n", result)
	},
}

func LongestSub (s string) string {
	ans := 0	
	strResult := map[string]int{}
	i := 0
	for j,rune := range s {
		char := fmt.Sprintf("%c", rune)
		if _, ok := strResult[char]; ok {
			i = int(math.Max(float64(strResult[char]),float64(i)))
		}
		ans = int(math.Max(float64(ans),float64(j-i+1)))
		strResult[char] = j+1 
	}
	result := fmt.Sprintf("result: %v",ans)
	return result
}


func init() {
	rootCmd.AddCommand(longestsubCmd)
	longestsubCmd.PersistentFlags().StringVarP(&inputstr, "string","s", "", "string")
	longestsubCmd.MarkFlagRequired("string")
}
