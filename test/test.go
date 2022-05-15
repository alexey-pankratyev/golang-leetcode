/*
Copyright © 2022 aleksei alexey.pankratev@gmail.com

*/
package main

import (
	"fmt"
)

func main() {
	arr0 := []string{"men","woman", "child", "woman"}
	map0 := map[int]string{}
	mapRes := map[string]int{}
	arrResult := []string{}
	for i:=0; i<len(arr0);i++ {
		map0[i] = arr0[i]
		fmt.Printf("%v\n",arr0[i])
	} 	
	for k, v := range map0 {
		if _, ok := mapRes[v]; ok {
			delete(map0,k)
		}else{
			mapRes[v] = k
			arrResult = append(arrResult,v) 	
		}
	}
	fmt.Printf("%v\n",arrResult)
}