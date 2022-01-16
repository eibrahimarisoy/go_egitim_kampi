package main

import (
	"fmt"
	"strings"
)

// import "fmt"

// var degisken_1 string
// var degisken_2 = "degisken_2"

// const sabit_1 = "deger 1"
// const (
// 	sabit_2 = "deger 2"
// 	sabit_3 = "deger 3"
// 	sabit_4 = "deger 4"
// )

// const (
// 	sabit_5 = iota
// 	sabit_6
// 	sabit_7
// )

// func main() {
// 	degisken_1 = "degisken_1"
// 	degisken_3 := "degisken_3"
// 	println(degisken_1, degisken_2, degisken_3)

// 	println(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6, sabit_7)
// }

// arrays

// var arr_1 [3]int
// var arr_2 = [5]int{1, 2, 3, 4, 5}

// func main() {
// 	arr_3 := make([]int, 3)

// 	arr_3[1] = 2

// 	fmt.Println(arr_1, arr_2, arr_3)
// 	fmt.Printf("%d", len(arr_1))
// }

// slices

// var slc_1 []int

// func main() {
// 	slc_2 := make([]int, 0, 3)
// 	// slc_2[0] = 2

// 	slc_2 = append(slc_2, 1)
// 	slc_2 = append(slc_2, 1)
// 	slc_2 = append(slc_2, 1)
// 	slc_2 = append(slc_2, 1)
// 	slc_2 = append(slc_2, 1)
// 	slc_2 = append(slc_2, 1)
// 	slc_2 = append(slc_2, 1)

// 	fmt.Println(slc_1, slc_2)
// 	fmt.Printf("slc_1 len:%d cap:%d", len(slc_1), cap(slc_1))
// 	fmt.Printf("slc_2 len:%d cap:%d", len(slc_2), cap(slc_2))
// }

//  maps
// var map_1 map[int]string

// func main() {
// 	map_1 = make(map[int]string)
// 	map_2 := make(map[int]int)
// 	map_3 := make(map[int]int, 3) // boyut

// 	map_1[0] = "1"
// 	map_2[0] = 2
// 	map_3[0] = 3
// 	map_3[1] = 3
// 	map_2[2] = 3
// 	map_2[3] = 3

// 	fmt.Println(map_1, map_2, map_3)
// 	fmt.Printf("map_1 len:%d \n", len(map_1))
// 	fmt.Printf("map_2 len:%d \n", len(map_2))
// 	fmt.Printf("map_3 len:%d \n", len(map_3))

// }

// string

func main() {
	str := "lorem ipsum dolor sit amet, consect"

	str_1 := str[:5]
	str_2 := str[len(str)-4:]
	str_3 := fmt.Sprintf("%s ipsum dolor sit %s", str_1, str_2)

	if strings.EqualFold(str_1, "LOrEM") {
		fmt.Println("str_1 equal to LOrEM")

	}

	fmt.Println(strings.HasPrefix(str, "lorem"))

	fmt.Println(str)
	fmt.Println(str_1)
	fmt.Println(str_2)
	fmt.Println(str_3)
	fmt.Println(strings.ToUpper(str))

}
