package main

/*func main() {
//age := []float64{12.4, 4.5, 5.8, 6.2, 7.0, 8.8}
//words := []string{"banana", "apple", "orange", "grape", "pineapple"}
//n := len(age)
//var temp int

//selection sort
/*for i := 0; i < n; i++ {
	for j := i + 1; j < n; j++ {
		if age[i] > age[j] {
			temp = age[i]
			age[i] = age[j]
			age[j] = temp
		}
	}
}*/

//function to sort
//sort.Ints(age)
/*sort.Float64s(age)
fmt.Println(age)
sort.Strings(words)
fmt.Println(words)*/

//count the occurence of an element

//find the first repeated
/*arr := [6]int{12, 34, 12, 56, 12, 56}
	var item int
	index := -1
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 6; j++ {
			if arr[i] == arr[j] {
				item = arr[j]
				index = j
				goto OUT
			}
		}
	}

OUT:
	if index != -1 {
		fmt.Printf("Number %d repeated at %d", item, index)
	} else {
		fmt.Printf("No number repeated")
	}
}*/

//insert an element
/*arr := []int{1, 2, 3, 4, 5} //slice
	newitem := 6
	index := 2
	newArray := insertIntoArray(arr, newitem, index)
	fmt.Println("Original array: ", arr)
	fmt.Println("New array: ", newArray)
}

func insertIntoArray(arr []int, item, index int) []int {
	//create a new slice  with increased length
	newArr := make([]int, len(arr)+1)

	//copy elements before insertion point
	copy(newArr[:index], arr[:index])

	//insert the new item
	newArr[index] = item

	//copy the elements after the insertion point
	copy(newArr[index+1:], arr[index:])
	return newArr
}*/

// 2d array
/*func main() {
	var matrix [3][3]int

	fmt.Println("Enter the elements of matrix: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	sum1 := 0
	sum2 := 0
	fmt.Println("Left diagonal is: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				fmt.Printf("%d ", matrix[i][j])
				sum1 = sum1 + matrix[i][j]
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n")
	}
	fmt.Println("Right diagonal is: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 3-1-i {
				fmt.Printf("%d ", matrix[i][j])
				sum2 = sum2 + matrix[i][j]
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

	//add two matrix
    for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i][j]=mat1[i][j]+mat2[i][j]
		}

	}

}*/

//fibonacci series using goto
//calculate the power of given number using goto
/*for i=0;i<n;i++{

}
*/
//3d array
//matrix no-row no-col no

//rev number using goto and for loop
/*while(num!=0){
	digit=num%10
	rev=rev*10+sum
	num/=10
}*/

//check palindrome using label and for loop
/*if(rev==temp){
	fmt.Printf("Yes")
}else{
	fmt.Printf("No")
}*/

//check armstrong number(for a 3 digit number)
/*while(num!=0){
	digit=num%10
	sum=sum+digit*digit*digit
	num/=10
}
if(temp==sum){
	fmt.Printf("Yes")
}else{
	fmt.Printf("No")
}*/

//check if it is a prime number
/*for i=2;i<n;i++{
	if(num%i==0){
		flag=1;
		break;
	}
}*/

//check if it is a perfect number(sum of factors is that number)
/*for i=1;i<n;i++{
	if(num%i==0){
		sum=sum+i;
	}
}
if(num==sum){
	fmt.Printf("Yes")
}else{
	fmt.Printf("No")
}*/

//linear and binary search
//linear search
/*for i=0;i<n;i++{
	if(a[i]==x)
	{
		index=i;
		break;
	}
}*/
//bubble sort(asc desc)
/*for i=0;i<n;i++{
	for j=0;j<n-i-1;j++{
		if(a[j]>a[j+1]){
			temp=a[j];
			a[j]=a[j+1]
			a[j+1]=temp
		}
	}
}*/
//insertion sort(asc desc)
//name,roll no,(paper1, paper2,paper3)2>40=pass, paper3,total,percentage
