package main

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

// func main() {
// 	str := "lorem ipsum dolor sit amet, consect"

// 	str_1 := str[:5]
// 	str_2 := str[len(str)-4:]
// 	str_3 := fmt.Sprintf("%s ipsum dolor sit %s", str_1, str_2)

// 	if strings.EqualFold(str_1, "LOrEM") {
// 		fmt.Println("str_1 equal to LOrEM")

// 	}

// 	fmt.Println(strings.HasPrefix(str, "lorem"))

// 	fmt.Println(str)
// 	fmt.Println(str_1)
// 	fmt.Println(str_2)
// 	fmt.Println(str_3)
// 	fmt.Println(strings.ToUpper(str))

// }

// loops

// func main() {
// slc_1 := []int{1, 2, 3}
// slc_2 := []int{}

// for i, value := range slc_1 {
// 	fmt.Printf("index: %d value %d \n", i, value)
// }

// for i := 0; i < 10; i++ {
// 	slc_2 = append(slc_2, i)
// }

// for i := range [10]int{} {
// 	slc_2 = append(slc_2, i)
// }
// fmt.Println(slc_2)

// for _, value_2 := range slc_2 {
// 	for _, value_1 := range slc_1 { //range olduğı için yeni bir değer yaratılır diziyi etkilemez
// 		value_1 += value_2
// 	}
// }

// for _, value_2 := range slc_2 {
// 	for i := range slc_2 {
// 		slc_1[i] += value_2
// 	}
// }

// fmt.Println(slc_1)
// fmt.Println(slc_2)
// fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
// fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))

// 	c := time.After(5 * time.Second)

// 	for {

// 		b := false
// 		select {
// 		case <-c:
// 			b = true
// 		default:
// 			fmt.Println(time.Now())
// 			time.Sleep(1 * time.Second)
// 		}
// 		if b {
// 			break
// 		}
// 	}
// }

//structs
// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"
// )

// type Kullanici struct {
// 	// Ad      string      `json:"adiiii"` // tag
// 	// Soyad   string      `json:"-"`
// 	Ad    string
// 	Soyad string

// 	Takipci []Kullanici `json="takipci,omitempty"`

// 	// Begeni []struct {
// 	// 	// Tarih time.Time
// 	// }
// 	// Begeni []Begen
// }

// type Begen struct {
// 	Tarih time.Time
// }

// func main() {
// 	k := Kullanici{
// 		Ad:    "Go",
// 		Soyad: "Türkiye",
// 		Takipci: []Kullanici{
// 			{
// 				Ad:    "Takipci",
// 				Soyad: "1",
// 			},
// 			{
// 				Ad:    "Takipci",
// 				Soyad: "2",
// 			},
// 		},
// 	}

// 	arr, _ := json.Marshal(k)
// 	fmt.Println(string(arr))
// 	fmt.Println(k)

// }

// pointers

// type String *string

// func main() {
// 	var rstr String
// 	var pstr string

// 	fmt.Println(rstr)
// 	fmt.Println(pstr)

// 	pstr = "go türkiye"
// 	rstr = &pstr

// 	fmt.Println(rstr)
// 	fmt.Println(pstr)

// 	pstr = "go türkiye eğitim kampı"

// 	fmt.Println(*rstr)
// 	fmt.Println(pstr)

// 	*rstr = "go türkiye eğitim kampı 2021"
// 	fmt.Println(*rstr)
// 	fmt.Println(pstr)

// 	// replaceStr(pstr)

// 	// fmt.Println(*rstr)
// 	// fmt.Println(pstr)

// 	replaceStr(rstr)

// 	fmt.Println(*rstr)
// 	fmt.Println(pstr)
// }

// // func replaceStr(s string) {
// // 	s = strings.Replace(s, "go", "GO", 1)
// // }

// func replaceStr(s String) {
// 	*s = strings.Replace(*s, "go", "GO", 1)
// }

// functions

// func timer(c <-chan time.Time) {
// 	for {
// 		select {
// 		case <-c:
// 			return
// 		default:
// 			fmt.Println(time.Now())
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func timer(c <-chan time.Time, message ...string) {
// 	for {
// 		select {
// 		case <-c:
// 			return
// 		default:
// 			fmt.Println(time.Now(), message)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	timer(time.After(5*time.Second), "Selam")

// }

// func timer(c <-chan time.Time, message ...string) {
// 	defer fmt.Println("bir sonraki aşamaya geçiliyor") // fonksiyon bitince ters yönde çalışır
// 	defer fmt.Println("timer bitti")
// 	for {
// 		select {
// 		case <-c:
// 			return
// 		default:
// 			fmt.Println(time.Now(), message)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func timer(c <-chan time.Time, message ...string) {
// 	defer fmt.Println("bir sonraki aşamaya geçiliyor") // fonksiyon bitince ters yönde çalışır
// 	defer fmt.Println("timer bitti")

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	for {
// 		select {
// 		case <-c:
// 			return
// 		default:
// 			fmt.Println(time.Now(), message)
// 			time.Sleep(1 * time.Second)

// 			if time.Now().Day() == 16 {
// 				panic("beklenmeyen bir hata oluştu") // çalışmay engeller
// 			}
// 		}
// 	}

// }

//receivers structa değişiklilk yapacaksak pointer receviers kullanmalıyız

// type Kullanici struct {
// 	Ad      string `json:"adi"`
// 	Soyad   string `json:"soyad"`
// 	Takipci []Kullanici
// }

// func (k Kullanici) TakipciSayisi() int {
// 	return len(k.Takipci)
// }

// func (k *Kullanici) TakipciEkle(t Kullanici) {
// 	fmt.Println(k.Takipci)
// 	if k.Takipci == nil {
// 		k.Takipci = []Kullanici{}
// 	}
// 	k.Takipci = append(k.Takipci, t)
// }

// func main() {
// 	k := Kullanici{
// 		Ad:    "Go",
// 		Soyad: "Türkiye",
// 		Takipci: []Kullanici{
// 			{
// 				Ad:    "Takipci",
// 				Soyad: "1",
// 			},
// 			{
// 				Ad:    "Takipci",
// 				Soyad: "2",
// 			},
// 		},
// 	}

// 	t := Kullanici{
// 		Ad:    "Takipci",
// 		Soyad: "3",
// 	}
// 	fmt.Println(k.TakipciSayisi())
// 	k.TakipciEkle(t)
// 	fmt.Println(k.TakipciSayisi())

// }

//errors

// type Kullanici struct {
// 	Ad      string `json:"adi"`
// 	Soyad   string `json:"soyad"`
// 	Takipci []Kullanici
// }

// func (k Kullanici) TakipciSayisi() int {
// 	return len(k.Takipci)
// }

// func (k *Kullanici) TakipciEkle(t Kullanici) {
// 	if k.TakipciSayisi() == 10 {
// 		return errors.New("maximum rakipçi sayısına ulaşıldı")
// 	}

// 	if k.Takipci == nil {
// 		k.Takipci = make([]Kullanici, 0, 10)
// 	}
// 	k.Takipci = append(k.Takipci, t)
// 	return nil
// }
