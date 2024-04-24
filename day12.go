package main

// assigning a func to a variable
/*func test(x string) {
	fmt.Println("Hello", x)
}
func main() {
	x := test
	x("Saloni")
	test("Saloni")
}*/

// assigning an anonymous(should be defined inside the function) func to a var

/*func main() {
	test := func(x string) {
		fmt.Println(x)
	}
	test("Saloni")
}*/

// Note: inferred var cannot be declared globally(same goes with the var to which a function is assigned)
/*var b = 20

func main() {
	a := 10
	fmt.Println(a, b)
}
*/

/*
	func main() {
		test := func(x string) string {
			return x + ",you are awesome."
		}("Saloni")
		fmt.Println(test)
	}
*/
/*func calc(x, y int) (int, int) {
	return x + y, x - y
}
func main() {
	sum, diff := calc(10, 20)
	fmt.Printf("Addition is %d and Subtraction is %d", sum, diff)
}*/

/*
	func main() {
		YYYY, MM, DD := time.Now().Date()
		fmt.Println("Today's date is", DD, "/", MM, "/", YYYY)
	}
*/

/*func main() {
	currentDateTime := time.Now()
	day := currentDateTime.Day()
	month := currentDateTime.Month()
	year := currentDateTime.Year()
	hour := currentDateTime.Hour()
	min := currentDateTime.Minute()
	sec := currentDateTime.Second()

	fmt.Printf("Report is printed on %d/%d/%d at %d:%d:%d", day, month, year, hour, min, sec)
}*/

/*func main() {
	fmt.Println("Hello World")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Hello India")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Hello, Welcome to the world of Golang!")
}*/

/*func main() {
	date1 := time.Date(2020, 2, 5, 10, 5, 20, 0, time.UTC)
	date2 := time.Date(2020, 2, 5, 10, 5, 20, 0, time.UTC)
	date3 := time.Date(2020, 2, 6, 10, 5, 20, 0, time.UTC)

	if date1.Equal(date2) == true {
		fmt.Println("Date 1 and Date 2 are equal")
	} else {
		fmt.Println("Date 1 and Date 2 are not equal")
	}

	if date2.Equal(date3) == true {
		fmt.Println("Date 2 and Date 3 are equal")
	} else {
		fmt.Println("Date 2 and Date 3 are not equal")
	}
}*/

// for loop
/*func main() {
	//var i int
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	num := 0

	//for as while
	for num < 5 {
		fmt.Println(num)
		num++
	}

	//for as do while
	n := 1
	for {
		fmt.Println(n)
		n++
		if n > 5 {
			break
		}
	}
}
*/
/*func main() {
	str := "Hello World"
	fmt.Println("ASCII values of characters of given string")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c-%d\n", str[i], str[i])
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
}*/
/*func main() {
	var num int
	fmt.Printf("Enter the number: ")
	fmt.Scan(&num)
	multiplier := 1
	for multiplier <= 10 {
		product := num * multiplier
		fmt.Printf("%d * %d = %d\n", num, multiplier, product)
		multiplier++
	}
}
*/

// infinite loop
/*func main() {
	for true {
		fmt.Println("Hi")
	}
}*/

//continue
//break

// slice
/*func main() {
	colors := []string{"Red", "Yellow", "Green", "Blue"} //[]-size dediya to array hojaaega
	for index, val := range colors {
		fmt.Println(index, "-", val)
	}

	for _, val := range colors {
		fmt.Println(val)
	}

	for index, _ := range colors {
		fmt.Println(index)
	}

}*/

//nested loop

// array
/*func main() {
	array1 := [5]int{1, 4, 6, 7, 8} //var array[5] int
	fmt.Println(array1)
	fmt.Printf("%T\n", array1)

	array2 := [5]string{"Red", "Violet", "Green", "Blue"}
	fmt.Println(array2)

	array3 := [5]byte{'c', 'd', 'g', 't'}
	fmt.Println(array3)
}*/

// initialize some particular values
/*func main() {
	array := [5]int{0: 7, 3: 9}      //array := [5]int--------all zeroes
	fmt.Println(array)
}*/

// changing the element
/*func main() {
	weather := [3]string{"Rainy", "Sunny", "Cloudy"}
	weather[2] = "Stormy"
	fmt.Println(weather)
}*/

/*func main() {
	var weather = [...]string{"Rainy", "Sunny", "Cloudy"} //undefined array size
	fmt.Println(weather)
	length := len(weather)
	fmt.Println("No. of elements: ", length)
}*/
