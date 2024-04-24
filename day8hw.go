
package main

import "fmt"

// function prototype
func main() {
   var n1 int=5
   res:=fact(n1)
   fmt.Println("The factorial of given number is",res)
}

func fact(n int) int {
   if(n>=1){
   return n*fact(n-1)
   }else{
   return 1
   }
}