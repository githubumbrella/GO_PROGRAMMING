//structure

package main

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

/*func main() {
	var B1, B2 Books
	B1.title = "Go programming"
	B1.author = "Saloni"
	B1.subject = "History"
	B1.book_id = 619649

	B2.title = "Come programming"
	B2.author = "Aman"
	B2.subject = "Geography"
	B2.book_id = 584961

	fmt.Println("Book 1 title: ", B1.title)
	fmt.Println("Book 1 author: ", B1.author)
	fmt.Println("Book 1 subject: ", B1.subject)
	fmt.Println("Book 1 book_id: ", B1.book_id)

	fmt.Println("Book 2 title: ", B2.title)
	fmt.Println("Book 2 author: ", B2.author)
	fmt.Println("Book 2 subject: ", B2.subject)
	fmt.Println("Book 2 book_id: ", B2.book_id)

}*/

// passing the structure as an argument
/*func main() {
	var B1, B2 Books
	B1.title = "Go programming"
	B1.author = "Saloni"
	B1.subject = "History"
	B1.book_id = 619649

	B2.title = "Come programming"
	B2.author = "Aman"
	B2.subject = "Geography"
	B2.book_id = 584961
	printbook(B1)
	printbook(B2)
}
func printbook(book Books) {
	fmt.Println("Book title: ", book.title)
	fmt.Println("Book author: ", book.author)
	fmt.Println("Book subject: ", book.subject)
	fmt.Println("Book book_id: ", book.book_id)
}
*/

// func inside a struct
// initialize the func Rectangle
/*type Rectangle func(int, int) int

// create a structure
type rectanglePara struct {
	length  int
	breadth int
	color   string

	//function as a field of struct
	rect Rectangle
}

func main() {
	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}
	fmt.Println("Color of rectangle: ", result.color)
	fmt.Println("Area of rectangle: ", result.rect(result.length, result.breadth))
}
*/

// slice
/*func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)
}*/

/*func main() {
	var n []int
	//make function to allocate the memory
	n = make([]int, 3, 5) //length and capacity(max number to which the length can be exceeded)
	n[2] = 6
	printSlice(n)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v", len(x), cap(x), x)
}*/

// slice from an array
/*func main() {
	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	sliceNumbers := numbers[4:7]
	fmt.Println(sliceNumbers)
}*/

// slice functions
// append(),copy(),Equal(),len
/*func main() {
	primeNumbers := []int{2, 3}

	primeNumbers = append(primeNumbers, 5, 7)

	fmt.Println(primeNumbers)
}*/

/*func main() {
	even := []int{2, 4}
	odd := []int{3, 5}

	even = append(even, odd...)
	fmt.Println(even)
}*/
