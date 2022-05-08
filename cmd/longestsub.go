/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

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

func LongestSub (s string) string{
	stringIter := ""
	maxCount := 0
	strResult := ""
	for pos,rune := range s {
		char := fmt.Sprintf("%c", rune)
		if strings.Contains(stringIter,char) {
			if maxCount < len(stringIter){
				maxCount = len(stringIter)
				strResult = stringIter	
			}
			i := strings.Index(stringIter, char)								
			stringIter = stringIter[i+1:]+char
		} else {
			stringIter = stringIter + char
			if pos == len(s)-1 && maxCount < len(stringIter){
				strResult = stringIter	
			}
		}	
	}
	result := fmt.Sprintf("result: str-\"%v\", len-\"%v\"",strResult,len(strResult))
	return result
}

func init() {
	rootCmd.AddCommand(longestsubCmd)
	longestsubCmd.PersistentFlags().StringVarP(&inputstr, "string","s", "", "string")
	longestsubCmd.MarkFlagRequired("string")
}
