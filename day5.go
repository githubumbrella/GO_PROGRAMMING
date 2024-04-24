/*func main(){
	num:=5
	fmt.Printf("%T %T\n",num,&num)
	fmt.Print(num,&num)
}*/

/*func main(){
	num,floatnum:=1238,5332.67489
    fmt.Printf("Default: %f\n",floatnum)
	fmt.Printf(".2f: %5.2f\n",floatnum)
	fmt.Printf(".4f: %.4f\n",floatnum)
    fmt.Printf("Scientific: %e\n",floatnum)
	fmt.Printf("Decimal %d\n",num)
	fmt.Printf("Binary %b\n",num)
	fmt.Printf("Octal %o\n",num)
	fmt.Printf("Hexadecimal %X\n",num)
}*/

/*func main(){
	//fmt.Printf("%*s\n",40,"text")   //prints (40-4) spaces and text
	//fmt.Printf("%040s","text")      //prints (40-4) 0 and text
	var chr byte='0'
	str:="ABC"
	fmt.Printf("%T %T\n",chr,str)
	fmt.Printf("Character value: %c\n",chr)
	fmt.Printf("ASCII value: %d\n",chr)
}*/

/*func main(){
	var val float64=-19.25
	fmt.Printf("Absolute value of %f is %f",val,math.Abs(val))      //Abs() function accepts only float number
}*/

/*func main(){
	var num1 float64=10.25
	var num2 float64=20.14
	var large,small float64
	large=math.Max(num1,num2)
	fmt.Printf("Largest number is %f\n",large)
	large=math.Max(78,34)
	fmt.Printf("Largest number is %f\n",large)

	small=math.Min(num1,num2)
	fmt.Printf("Largest number is %f\n",small)
	small=math.Min(78,34)
	fmt.Printf("Largest number is %f\n",small)

	var num float64=10.2
	var p float64=3.0
	var result float64
	result=math.Pow(num,p)
	//result=math.Pow(10,3)
	fmt.Printf("Result: %f",result)
}*/

/*func main(){
	n1:=10.25
	fmt.Printf("%f\n",math.Ceil(n1))      //round up to the nearest integer(bigger one)
	fmt.Printf("%f\n",math.Floor(n1))     //round up to the nearest integer(lower one)
	println("Hello")                      //Use println instead of fmt.Println
	print("Hello")                        //Use print instead of fmt.Print
}*/

/*func main(){
	x:=100
	y:=10
	fmt.Printf("Addition: %d\n",x+y)
	fmt.Printf("Subtraction: %d\n",x-y)
	fmt.Printf("Multiplication: %d\n",x*y)
	fmt.Printf("Division: %d\n",x/y)
	fmt.Printf("Modulo: %d\n",x%y)
	x++
	fmt.Println("After post increment: ",x)
	x+=7        //x=x+7
	fmt.Println(x)
	x-=7        //x=x-7
	fmt.Println(x)
	x*=7        //x=x*7
	fmt.Println(x)
	x/=7        //x=x/7
	fmt.Println(x)
}
*/
/*
func main(){
	/*var a,b int
	fmt.Println("Enter two number: ")
    fmt.Scanln(&a,&b)
	var r1,r2 int
	r1=a%2
	r2=b%2
	fmt.Println(r1,r2)*/

/*x:=5
	y:=3
	fmt.Println(x&y)       // bitwise and
	fmt.Println(x|y)       //bitwise or
	fmt.Println(x^y)       //bitwise xor
	fmt.Println(x>>2)      //right shift (value decreases)
	fmt.Println(x<<2)      //left shift  (value increases)
	fmt.Println(x==y,x!=y,x>y,x<y,x>=y,x<=y)     //comparison operators

}
*/
/*package main
import(
	"fmt"
	//"unsafe"
	"reflect"
	//"math"
)*/
/*func main(){
	var num1 int=10
	var num2 byte=20
	var num3 float64=10.2
	//var name string="Saloni"

	fmt.Printf("Size of num1 is %d\n",unsafe.Sizeof(num1))
	fmt.Printf("Size of num2 is %d\n",unsafe.Sizeof(num2))
	fmt.Printf("Size of num3 is %v\n",unsafe.Sizeof(num3))
	//fmt.Printf("Size of name is %s",unsafe.Sizeof("Saloni"))
}*/
/*func main(){
	a:=10
	b:=10.2
	c:="Hello"
	d:=true
	e:=[]string{"India","Pakistan"}

	fmt.Println("Type of a is ",reflect.TypeOf(a))
	fmt.Println("Type of b is ",reflect.TypeOf(b))
	fmt.Println("Type of c is ",reflect.TypeOf(c))
	fmt.Println("Type of d is ",reflect.TypeOf(d))
	fmt.Println("Type of e is ",reflect.TypeOf(e))


	//another method used for same purpose
	fmt.Println("Type of a is ",reflect.ValueOf(a).Kind())
	fmt.Println("Type of b is ",reflect.ValueOf(b).Kind())
	fmt.Println("Type of c is ",reflect.ValueOf(c).Kind())
	fmt.Println("Type of d is ",reflect.ValueOf(d).Kind())
	fmt.Println("Type of e is ",reflect.ValueOf(e).Kind())
}*/

package main

//defer
/*func main(){
	defer fmt.Println("Hello\n")
	fmt.Println("Good morning\n")
	defer fmt.Println("Hii\n")
	fmt.Println("Good evening\n")
}*/

/*func main(){
	var radius float32
	var area,perimeter float32
    const PI=3.14
	fmt.Println("Enter radius: ")
	fmt.Scanln(&radius)
	area=PI*radius*radius
	perimeter=2*PI*radius
	fmt.Printf("Area and perimeter of circle are: %.2f and %.2f",area,perimeter)
}*/

/*func main(){
	var ftemp float64
	var ctemp float64

	fmt.Printf("Enter the temperature in Fahrenheit: ")
	fmt.Scanf("%f",&ftemp)
	ctemp=(ftemp-32)/1.8
	fmt.Printf("Temperature in Celcius: %.2f",ctemp);
	ftemp=(ctemp*1.8)+32;
	fmt.Println("Temperature in Fahranheit: %.2f",ftemp);
}*/
