/*package main
import "fmt"

func main(){
	var a,b,c,d =1,3.7,5,7.9
	u,v :=7, "Saloni"
	fmt.Println(a,b,c,d)
	fmt.Println(u,v)
}*/

/*func main(){
	var(
		b int
		a float32 = 1.1
		c string="Saloni"
	)
	fmt.Println(b,a,c)
	fmt.Printf("%d %v %s",b,a,c)
}*/

/*a:=10; var b int=5
func main(){
	fmt.Println(a,b)
}----":=" cannot be declared outside the function but var can be declared*/

//:= doesn't permit the separate declaration and initialization

//Scanf function
/*func main(){
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your surname: ")
	var fullname string
	fmt.Scanln(&fullname)
	fmt.Println(name+fullname)   //Addition of two strings-Concatenation

	fmt.Print("saloni");         //No new line is generated
}*/

/*func main(){
	var name string="Saloni Vishwakarma"
	var age int=21
	var branch string="Cyber Security"
    var collegename string="RCOEM"
	fmt.Println("Name: ",name,"Age: ",age," \nBranch: ",branch,"  College name: ",collegename)
}*/

/*func main(){
	a:="Saloni Vishwakarma"
	fmt.Println("My name is","'",a,"'")
	fmt.Printf("My name is '%s'",a)
}
*/
/*func main(){
	const(
		name="GeeksForGeeks"
		dept="CS"
	)
	fmt.Printf("%s is a %s portal",name,dept)
}*/

/*func main(){
	var num3 int = 15
	fmt.Printf("The binary value of %d is %b\n",num3,num3)
	var num4 float32 = 15.897
	fmt.Printf("The floating point is %g\n",num4)
	var num5 = 15 +2i
	fmt.Printf("Scientific notation is %e\n",num5)
}*/

//Universal format specifier
/*func main(){
	var a int=10
	var b float32=87.09
	var name string="Saloni"
	fmt.Printf("%v %v %v",a,b,name)
}*/

//lvalues(refers to a memory loc) and rvalues(data value stored at that loc)
//b=10
//literals are always rvalues

//Integer value literals
/*
0x 0X-HEX
0 0o 0O-OCTAL
0B 0b-BINARY
NO PREFIX -DECIMAL
*/

/*func main(){
	fmt.Println(15==017)
	fmt.Println(15==0XF)
	a:=6
    a++
	fmt.Println(a);
}
//pre incremnent/decrement operators don't work in go
*/
/*func main(){
	a:=15;
	fmt.Printf("Binary of %d is %0b\n",a,a)
	fmt.Printf("Hexadecimal of %d is %0X\n",a,a)
	fmt.Printf("Octal of %d is %0o\n",a,a)
	c:=1.23e2
	fmt.Println(c)
}*/
//Format specifiers
/*v-default format
d-decimal
g-float
b-binary
o-octal
t-boolean value
s-string
*/
package main

import (
	"fmt" ////to import multiple libraries
)

/*func main(){
//fmt.Printf("Expression 15==0xF is %t\n",15==0xF)------bool
//fmt.Printf("Expression 15==0xF is %T",15==0xF)--------datatype
/*fname:="Saloni"
lname:="Vishwakarma"
fmt.Print("My name is ",fname," and my last name is ",lname,"\n")
i:="Stupid"
j:="Cupid"
fmt.Print(i,"\n")
fmt.Print(j,"\n")
fmt.Printf("%s\t%s\n",i,j)
a,b:=10," Pointer"
fmt.Print(a,b)*/
/*var i=15.5
	var txt="Hello World!"
	fmt.Printf("%v\n",i)
	fmt.Printf("%#v\n",i)
	fmt.Printf("%v%%\n",i)  //prints the % once if written twice
	fmt.Printf("%T\n",i)
	fmt.Printf("%v\n",txt)
	fmt.Printf("%#v\n",txt) //# for actual format
    fmt.Printf("%T\n",txt)
	fmt.Printf("%v\n",txt)*/
/*chr:='a'
fmt.Printf("%c",chr)*/

/*var i=152
	fmt.Printf("%b\n",i)        //binary
	fmt.Printf("%d\n",i)        //decimal
	fmt.Printf("%+d\n",i)       //sign of number
	fmt.Printf("%o\n",i)        //octal
	fmt.Printf("%O\n",i)        //Octal
	fmt.Printf("%x\n",i)        //Hex small
	fmt.Printf("%X\n",i)        //Hex capital
	fmt.Printf("%#x\n",i)       //actual format of hexadecimal number
	fmt.Printf("%4d\n",i)       //4 unit spaces are allocated -2 extra space
	fmt.Printf("%-4d\n",i)      //lefyt justified
	fmt.Printf("%04d\n",i)      //fill up the empty literals with 0


	var txt="Hello";
	fmt.Printf("%s\n",txt)
	fmt.Printf("%q\n",txt)
	fmt.Printf("%8s\n",txt)
	fmt.Printf("%-8s\n",txt)
	fmt.Printf("%x\n",txt)
	fmt.Printf("% x\n",txt)
}*/

func main() {
	name := "Saloni"
	rollno := 13
	var p1, p2, p3 int
	fmt.Printf("Enter the mark of paper 1 : ")
	fmt.Scanln(&p1)
	fmt.Printf("Enter the mark of paper 2 : ")
	fmt.Scanln(&p2)
	fmt.Printf("Enter the mark of paper 3 : ")
	fmt.Scanln(&p3)
	var sum int
	sum = p1 + p2 + p3
	fmt.Printf("Name of student: %s\t\t", name)
	fmt.Printf("Roll no: %d\n", rollno)
	fmt.Printf("Total marks obtained out of 300: %d\n", sum)
	fmt.Printf("Percentage obtained: %d%%", sum/3)
}
