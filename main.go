package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alexey-pankratyev/golang-leetcode/packages/twosum"
)

var twosumArr = []int{1,2,6,7,5,20}
var twosumTarget = 22 

func main() {
    cmd := &cobra.Command{
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("We run the TWOSUM module and the result is here:")
            fmt.Println(twosum.TwoSum(twosumArr, twosumTarget))
        },
    }
	
	fmt.Println("Calling cmd.Execute()!")
    cmd.Execute()
}