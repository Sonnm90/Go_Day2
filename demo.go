package main

import "fmt"

func main() {
	// Array
	//var a [3]int
	//a[0] = 1
	//a[1] = 2
	//a[2] = 3
	// b:= [3]int{1,2,3}
	//fmt.Println(a)
	////so sánh 2 array
	//fmt.Println(a==b)

	x := [...]int{1, 2, 3, 4, 5}
	y := x   // b là một array mới có giá trị giống a
	y[0] = 9 // Thay đổi giá trị một phần tử của b

	fmt.Println("a is ", x) // Kết quả: 1 2 3 4 5
	fmt.Println("b is ", y) // Kết quả: 9 2 3 4 5

	var a = [...]int{1, 2, 3}
	// b là một con trỏ tới array a
	var b = &a
	// in ra hai phần tử đầu tiên của array a
	fmt.Println(a[0], a[1])
	// truy xuất các phần tử của con trỏ array cũng giống như truy xuất các phần tử của array
	fmt.Println(b[0], b[1])
	// duyệt qua các phần tử trong con trỏ array, giống như duyệt qua array
	for index, value := range b {
		// thay đổi từng phần tử trong b
		fmt.Println(b)
		b[index] += 1
		fmt.Println(index, b[index], value)
	}
	for i := 0; i < len(a); i++ {
		a[i] += 1
		fmt.Println(i, a[i])
	}
	fmt.Println(b)
	//giá trị của các phần tử trong a bị thay đổi vì b
	for index, value := range a {
		fmt.Println(index, value)
	}

	arr1 := [2][2]int{
		{2, 4},
		{6, 8},
	}
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Println(arr1[i][j])
		}
	}
	for i, _ := range arr1 {
		for j, _ := range arr1[i] {
			fmt.Println(arr1[i][j])
		}
	}

	var e [0]int
	fmt.Println(e)

	//	String
	var s = "hello world"
	hello := s[:5]
	fmt.Println(hello)

	//Slices
	//var s1 []int
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))
	//s1 = append(s1, 1, 2, 3)
	//fmt.Println(s1)
	//var s2 = []int{1, 2, 3}
	//fmt.Println(len(s2))
	//fmt.Println(cap(s2))
	//
	//for i := range s1 {
	//	fmt.Println(s1[i])
	//}

	//old := []string{"r", "o", "a", "d"}
	//new := old[2:]
	//fmt.Println(new)
	//new[1] = "m"
	//fmt.Println(new)
	//fmt.Println(old)

	//sl := make([]int, 1)
	//printSlice(sl)
	//for i := 1; i < 5; i++ {
	//	sl = myAppend(sl, i)
	//}

	//arr2 := [6]int{0, 1, 2, 3, 4, 5}
	//s2 = arr2[2:4]
	//fmt.Println(s2)
	//fmt.Println(len(s2))
	//fmt.Println(cap(s2))
	//s2 = append([]int{5}, s2...)
	//fmt.Println(s2)
	//fmt.Println(len(s2))
	//fmt.Println(cap(s2))
	//fmt.Println(arr2)

	//var s3 = []int{0, 1, 2, 3, 4, 5}
	//fmt.Println(s3)
	//s3 = append(s3[:2], s3[3:]...)
	//fmt.Println(s3)
	//fmt.Println(copy(s3[2:], s3[3:]))
	//fmt.Println(s3)
	//s3 = s3[:2+copy(s3[2:], s3[3:])]
	//fmt.Println(s3)

	//Map
	var ageOfStudents = map[string]int{
		"Sơn":   33,
		"Chuân": 31,
	}
	ageOfStudents["Vượng"] = 30
	fmt.Println(ageOfStudents)
	name := "Sơn"
	fmt.Println(name, ageOfStudents[name])
	personSalary := map[string]int{}
	fmt.Println(personSalary)

	delete(ageOfStudents, name)
	value, check := ageOfStudents["Chuân"]
	fmt.Println(value, check)
	value1, check1 := ageOfStudents["Sơn"]
	fmt.Println(value1, check1)

	newAgeOfStudents := ageOfStudents
	newAgeOfStudents["Hà"] = 19
	fmt.Println(ageOfStudents)
	fmt.Println(newAgeOfStudents == nil)

	//Range
	letters := []string{"a", "b", "c"}

	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}

	//Without index and value. Just print array values
	fmt.Println("\nWithout Index and Value")
	i := 0
	for range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letters[i])
		i++
	}
}

func myAppend(sl []int, val int) []int {
	sl = append(sl, val)
	printSlice(sl)
	return sl
}

func printSlice(sl []int) {
	fmt.Printf("Slice %v\n", sl)
	fmt.Printf("Len %v, Cap %v\n\n", len(sl), cap(sl))
}
