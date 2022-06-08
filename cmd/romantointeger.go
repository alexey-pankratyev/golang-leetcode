/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	romantointegerin string
)

// romantointegerCmd represents the lromantointeger command
var romantointegerCmd = &cobra.Command{
	Use:   "romantointeger -s examplestring",
	Short: "Implementation of the problem of Roman to Integer for litcode",
	Long: `https://leetcode.com/problems/roman-to-integer
Given an roman, convert it to integer a numeral.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := romantointegerSub(romantointegerin)
		fmt.Printf("%v \n", result)
	},
}

func romantointegerSub (s string) int {
	romanToIntMap := map[string]int{ "M":1000, "D":500,"C":100, "L":50, "X":10, "V":5, "I": 1 }
	resInt:=0
	for i:=0; i<len(s); i++  {
		fmt.Print(i+1)
		if i+1 < len(s) && romanToIntMap[string(s[i])] < romanToIntMap[string(s[i+1])] {			
			resInt+=romanToIntMap[string(s[i+1])]-romanToIntMap[string(s[i])]
			fmt.Printf("r1: %v, r2: %v\n",romanToIntMap[string(s[i+1])],romanToIntMap[string(s[i])])
			i+=1
		}else {
			resInt+=romanToIntMap[string(s[i])]	
		}
		fmt.Printf("resInt: %v\n",resInt)
	}
	return resInt
}

func init() {
	rootCmd.AddCommand(romantointegerCmd)
	romantointegerCmd.PersistentFlags().StringVarP(&romantointegerin, "string","s", "", "string")
	romantointegerCmd.MarkFlagRequired("string")
}
