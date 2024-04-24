package main
import "fmt"

//constants(global variables)
/*const PI=3.14
func main(){
   PI=9
   fmt.Println(PI)
}*/

//Typed constants 
/*const A int =1
func main(){
    fmt.Println(A)
}*/

//Untyped constants: compiler itself decides the type of variable
/*const A =1
func main(){
    fmt.Println(A)
}*/

/*const(
  A int = 1
  B = 3.14
  C = "Hi!"
)
func main(){
   fmt.Println(A,B,C)
}*/

/*func main(){
  const l int =10
  const b int =20
  var area int
  area=l*b
  fmt.Printf("Area=%d",area)
}*/

/*func main(){
   var s1 string="John"  //statically declared type
   var s2="Jane"         //type is inferred 
   x:=2                  //type is inferred
   var b float64 = 20.0
   fmt.Printf("%d\n",x)
   fmt.Printf("%s\n",s1)
   fmt.Printf("%T\n",s2)   //%T for printing the type of variable
   fmt.Printf("%T\n",b)
}
*/

func main(){
   var a string
   var b int
   var c bool=true
   //var a,b,c = "saloni",34,true
   
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
}





