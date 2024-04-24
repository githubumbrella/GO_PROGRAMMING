/*package main
import(
"fmt"
"os"
)
func main(){
   var n1 string
   var n2 int
   if _,err=fmt.Scan(&name,&age);err!=nil{
      fmt.Println(err)
      os.Exit(1)
   }
   fmt.Printf("Your name is: %s\n",name)
   fmt.Printf("Your age is %d\n",age)
}

*/


/*package main
import(
"fmt"
)
func main(){
 var n1,n2 int
 fmt.Println("Enter 2 numbers: ")
 fmt.Scan(&n1,&n2)
 if(n1<n2){
     fmt.Printf("%d is the less than %d",n1,n2)
 }else{
     fmt.Printf("%d is the greater than or equal to %d",n1,n2)
 }
}*/

package main
import "fmt"
func main(){
  var n1 int
  fmt.Printf("Enter a number: ")
  fmt.Scan(&n1)
  if(n1%2==0){
     fmt.Printf("%d is even\n",n1)
  }else{
     fmt.Printf("%d is odd\n",n1)
  }

  if(n1<0){
      fmt.Printf("Absolute value of %d is %d",n1,n1*(-1))
  }else{
      fmt.Printf("Absolute value of %d is %d",n1,n1)
  }
} 