// day 10 24-02-2024
// Print Odd numbers using Golang Closure
/** func calculate() func() int{
     num:=1
	 // returns inner function
	 return func() int{
	    num=num+2
		return num
	}
}

func main(){
     // call the outer function
	 odd:=calculate()
	 // call the inner function
	 fmt.Println(odd())
	 fmt.Println(odd())
	 fmt.Println(odd())

	 // call the outer function again
	 odd2:=calculate()
	 fmt.Println(odd2())
	 } **/

// Go Closures - closure is a nested function that helps to access the outer function's variables even after the outer function is closed
//outer function
/** func greet() func() string{
     // variable defined outside the inner function
	 name:="John"
	 //return a nested anonymous function
	 return func() string{
	     name="Hi "+name
		 return name
	}
}

func main(){
     // call the outer function
	 message:=greet()
	 // call the inner function
     fmt.Println(message())
}	**/

// Closure helps in data isolation
/** func displayNumbers() func() int{
     number:=0
     // inner function
     return func() int{
        number++
     return number
     }
}

func main(){
     //returns a closure
     num1:=displayNumbers()
     fmt.Println(num1())
     fmt.Println(num1())
     fmt.Println(num1())
     // returns a new closure
	 num2:=displayNumbers()
	 fmt.Println(num2())
	 fmt.Println(num2())
} **/

// Fibonacci series
func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}