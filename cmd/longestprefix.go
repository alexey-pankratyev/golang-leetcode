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
	inputPrefixStr []string
)

// longestprefixCmd represents the longestprefix command
var longestprefixCmd = &cobra.Command{
	Use:   "longestprefix",
	Short: `longestprefix -s prefix1234,prefix345,prefix999`,
	Long: `Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string. 
For example:
Input: strs = ["flower","flow","flight"]
Output: "fl".`,
	Run: func(cmd *cobra.Command, args []string) {
		result := longestPrefix(inputPrefixStr)
		if len(result) > 0 {
			fmt.Println(result)
		} else {
			fmt.Println("Prefix wasn't find!")
		}		
	},
}

func longestPrefix(strs []string) string{
	var resPrefix string
	if len(strs)==1{
		resPrefix = strs[0]
	} else{
		resPrefix = "" 
	}		
	for iWord, world := range(strs){	
		prefix := ""	
		if iWord < len(strs)-1 {			
			for _, chr := range(world) {
				prefixCheck := prefix + string(chr)
				if strings.HasPrefix(string(strs[iWord+1]), prefixCheck){
					prefix = prefix + string(chr)
				}else{
					break
				}		
			}			
		}
		if prefix != "" {
			if len(resPrefix) == 0 {
				resPrefix = prefix
			}
			if len(resPrefix) > len(prefix) {
				resPrefix = prefix
			}
		}
		
	}
	if len(strs)!=1 {
		for _, world := range(strs){
			if ! strings.HasPrefix(world, resPrefix){
				resPrefix = ""
			}
		}
	}
	return resPrefix
}

func init() {
	rootCmd.AddCommand(longestprefixCmd)
    longestprefixCmd.Flags().StringSliceVarP(
        &inputPrefixStr, "prefixes", "s", inputPrefixStr, 
        "List of words with prefixes",
    )
	longestprefixCmd.MarkFlagRequired("prefixes")
}