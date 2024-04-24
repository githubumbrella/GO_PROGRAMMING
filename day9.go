package main
import "fmt"
/*func sum(number int)int {
     if number<=0{
           return 0
     }else{
           return number+sum(number-1)
     }
}

func main(){
     var num int 
     fmt.Println("Enter a number: ")
     fmt.Scan(&num)
     if(num<0){
          fmt.Println("Abe baitad, positive number daal")
     }else{
          var result=sum(num)
          fmt.Println("Sum:",result)
     }
}*/

//anonymous function
/*func main(){
   var greet=func(){
       fmt.Println("Hello, how are you")
   }
   welcome:=greet
   
 greet()
 welcome()
}*/

/*func main(){
   //anonymous func with parameter
   var sum=func(n1,n2 int)int{
       add:=n1+n2
       return add
   }
   result:=sum(5,3)
   fmt.Println("Sum is:",result)
}*/

//anonymous func as an argument
//var sum=0
func findSquare(num int)int{
    square:=num*num
    return square
}
func main(){
    sum:= func(n1 int,n2 int)int{
        return n1+n2
    }
    diff:= func(n1 int,n2 int)int{
        return n1-n2
    }
    mul:= func(n1 int,n2 int)int{
        return n1*n2
    }
    div:= func(n1 int,n2 int)int{
        return n1/n2
    }

    r1:=findSquare(sum(9,3))
    r2:=findSquare(diff(9,3))
r3:=findSquare(mul(9,3))
r4:=findSquare(div(9,3))

fmt.Println("Result is: ",r1)
fmt.Println("Result is: ",r2)
fmt.Println("Result is: ",r3)
fmt.Println("Result is: ",r4)

}

