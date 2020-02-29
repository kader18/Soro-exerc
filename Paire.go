package main

import ."fmt"

func kader(tab []int) []int{
     
	var result []int 
     for _,v:= range tab{
        
		if v%2 == 0 {
		  result = append(result, v)  
		}
     }
   return result 
}

func main() {
   Println(kader([]int{-1, 300, -300, 200, 10, 0, 5}))
}
