package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*func main() {
	student := make(map[int]string)
	student[1] = "saloni"
	student[2] = "madhavi"
	fmt.Println(student)
}
*/

/*func main() {
	places := map[string]string{"Nepal": "Kathamandu", "US": "Washington DC", "Norway": "Oslo"}

	for country := range places {
		fmt.Println(country)
	}
}*/

// structure used in the map
/*type Vertex struct {
	lat, long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.9832, -74.9285},
	"Google":    Vertex{35.097, -28.439287},
}

func main() {
	fmt.Println(m)
	m["Nagpur"] = Vertex{90.398, -98.389}
	fmt.Println(m)
}
*/

//test if a key is present in map wiht a two-value assignment

/*func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value :", v, "Present?", ok)

}*/

// pointer
/*func main() {
	var i = 20
	var ipointer *int
	ipointer = &i

	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(ipointer)

	*ipointer = 30
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(ipointer)
}
*/

/*func update(num *int) {
	*num = 30
}
func main() {
	var number = 55
	update(&number)
	fmt.Println(number)
}*/

//swapping using pointer
/*func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}
*/

//swapping without using pointer or 3rd var
/*func swap(x int, y int) (int, int) {
	return y, x
}
func main() {
	a := 10
	b := 20
	a, b = swap(a, b)
	fmt.Println(a, b)
}*/

//calling a web api
/*import(
	"fmt"
	"net/http"
	"io/ioutil"
)*/
func main() {
	//URL of the API endpoint
	apiUrl := "https://api.example.com/data"

	//Send GET request
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error making GET request: %s", err)
		return
	}

	defer response.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: \n", err)
		return
	}

	fmt.Println(string(body))
}
