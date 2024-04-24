// day 11 02-03-2024
// Golong program to check given year is leap year or not
/** func main(){
     var year int=0
     fmt.Printf("Enter number : ")
     fmt.Scanf("%d",&year)

     if (year>0){
     if(year%4==0 && year%100!=0) || (year%4==0 && year%100==0 && year%400==0){
            fmt.Printf("Entered year is a leap year")
     } else {
            fmt.Printf("Entered year is not a leap year")
     }
	 } else {
	        fmt.Printf("Invalid Year")
     }
}	**/

// nested if else statements (else if)
/** func main(){
     x:=8
	 y:=10
	 if(x!=y){
	    if(x<y){
		   fmt.Println("x is less than y");
		} else if(x>y){
		   fmt.Println("x is greater than y");
		}
	  } else {
	       fmt.Println("x is equal to y");
		 }
	} **/

/** func main(){
     var time int
	 fmt.Println("Enter time : ")
	 fmt.Scan(&time)
	 if(time>0 && time<=24){
	      if(time<=10){
		     fmt.Println("Good morning!")
		  } else if(time>10 && time<=20) {
		     fmt.Println("Good Day!")
		  } else if(time>20 && time<=24) {
		     fmt.Println("Good evening!")
		  }
	 } else {
	     fmt.Println("Invalid Time")
	}
} **/

// Switch Statements
/** func main(){
     thisMonth:=5
	 switch thisMonth{
	    case 1:
		    fmt.Println("January")
		case 2:
		    fmt.Println("February")
		case 3:
		    fmt.Println("March")
		case 4:
		    fmt.Println("April")
		case 5:
		    fmt.Println("May")
		case 6:
		    fmt.Println("June")
	}
} **/

// different message for particular day
/** func main(){
     today:=time.Now()
	 switch today.Day(){
	 case 5:
	    fmt.Println("Today is 5th.Clean your house.")
	 case 10 :
	    fmt.Println("Today is 10th.Buy some wine")
	 case 15 :
	    fmt.Println("Today is 15th.Visit a doctor")
	 case 25:
	    fmt.Println("Today is 25th.Buy some food")
	 case 31:
	    fmt.Println("Party")
	default: fmt.Println("No information available on this day")
	}
	} **/

// Switch with Multiple Case Values
/** func main (){
    thisMonth:=6
	switch thisMonth{
	    case 1,3,5,7,8,10,12:
		    fmt.Println("31 days")
		case 4,6,9,11:
		    fmt.Println("30 days")
		case 2:
		    fmt.Println("28 or 29 days")
	}
} **/

// Fallthrough Case Statement
// The fallthrough keyword in the case statement transfers the execution to the next case
/** func main(){
    x:=3
	switch x{
	    case 1:
		    fmt.Println("1")
			fallthrough
		case 3:
		    fmt.Println("3")
			fallthrough
		case 5:
		    fmt.Println("5")
		}
	} **/

// Type Switch
/** func main(){
    var x interface{}
	// var x interface{}="RKNEC"
	switch i:=x.(type){
	case nil:
	    fmt.Printf("Type of x: %T",i)
	case int:
	    fmt.Printf("x is int")
	case float64:
	    fmt.Printf("x is float64")
	case func(int) float64:
	    fmt.Printf("x is func(int)")
	case bool,string:
	    fmt.Printf("x is bool or string")
	default:
	    fmt.Printf("don't know the type")
	}
} **/

// Goto statement - allows you to transfer control to another part of the code within the same function
/** func main(){
    i:=0
	loop: // Label for the goto statement
	fmt.Println(i)
	i++
	if i<5{
	    goto loop // Jump to the loop label
	}
	fmt.Println("Loop ends here")
} **/

// Golang code to check a string contains a specified substring
/** func main(){
    var str string="Hello World"
	var subStr string="Wor"

	if strings.Contains(str,subStr) == true{
	    fmt.Printf("String (%s) contains sub-string (%s)",str,subStr)
	} else {
	    fmt.Printf("String (%s) does not contains substring (%s)",str,subStr)
	}
} **/

// Convert string in uppercase and lowercase in Golang
/** func main(){
    var str string="Hello World"
	var res1,res2 string
	res2=strings.ToUpper(str)
	res1=strings.ToLower(str)

	fmt.Println(res1,res2)
	} **/

// Getting the index of a character in the string
func main() {
	var str string = "Hello World"
	var ind int = 0
	ind = strings.Index(str, "W")
	fmt.Println("Index is : ", ind)
}