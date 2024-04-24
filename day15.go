package main

/*func main() {
	pn := []int{2, 3, 5, 7}
	n := []int{1, 2, 3}

	fmt.Println(cap(pn), cap(n))

	copy(n, pn)
	fmt.Println(n)

	fmt.Println(cap(pn), cap(n))

	n = append(n, 7, 8, 9, 0, 5, 0, 9, 7, 6, 8, 99, 88, 77)

	fmt.Println(cap(pn), cap(n))
}*/

/*func main() {
	arr := [10]int{10, 20, 40, 60, 50, 60, 70, 80, 90, 100}

	intSlice := arr[2:5]

	for index, ele := range intSlice {
		fmt.Println("Index: ", index, ",element: ", ele)
	}

	for _, ele := range intSlice {
		fmt.Println("Element: ", ele)
	}
	for index, _ := range intSlice {
		fmt.Println("Index: ", index)
	}
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	slice := []string{"Saloni", "Aman"}
	sort.Strings(slice)
	fmt.Println(slice)
}*

//check whether the slice is sorted or not
func main() {
	var status bool

	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{30, 10, 80, 70, 20}

	status = sort.IntsAreSorted(s1)

	if status == true {
		fmt.Println("Slice 1 is sorted")
	} else {
		fmt.Println("Slice 1 is not sorted")
	}

	//status = false
	status = sort.IntsAreSorted(s2)

	if status == true {
		fmt.Println("Slice 2 is sorted")
	} else {
		fmt.Println("Slice 2 is not sorted")
	}

}*/

// not using pre defined function
/*func main() {
	s2 := []int{30, 10, 80, 70, 20}
	var flag int = 0
	for i := 0; i < len(s2); i++ {
		for j := i + 1; j < len(s2); j++ {
			if s2[i] > s2[j] {
				flag = 1
				break
			}
		}
	}
	if flag == 1 {
		fmt.Println("Slice is not sorted")
	} else {
		fmt.Println("Slice is sorted")
	}
}*/

// map(key-value pair) in golang
/*func main() {
	subjectMarks := map[string]int{"Saloni": 13, "Aman": 31}
	fmt.Println(subjectMarks)
	fmt.Println(subjectMarks["Saloni"]) //accessing through the key
	fmt.Println(subjectMarks["Aman"])

	subjectMarks["Saloni"] = 1133 //updating the value
	fmt.Println(subjectMarks)
}*/

// appending and deleting the map variable with new key-value pair
/*func main() {
	students := map[int]string{1: "Saloni", 2: "Aman"}
	fmt.Println(students) //ascending order

	students[3] = "JCW"
	fmt.Println(students)

	students[10] = "KSH"
	fmt.Println(students)

	delete(students, 2)
	fmt.Println(students)

	for key, item := range students {
		fmt.Printf("%s is on number %d\n", item, key) //random order
	}
}
*/
/*func main() {

}*/
