package main

import ."fmt"

func kader(tab []int) []int{
  
  
   for i:=0; i<len(tab); i++ {
     for j,v:= range tab{
	   if tab[i] < v {
	      tab[j], tab[i] = tab[i], tab[j]
	   }
	 }
   }
   return tab
   
   
}

func main() {
   Println(kader([]int{-1, 200, -300, 200, 10}))
}
