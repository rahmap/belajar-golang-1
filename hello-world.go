package main

import (
	"fmt"
	"strings"
)

func main() {
	//name := "Rahma"
	//name = "Feby"
	//var age uint8 = 19
	//
	//address := "Turi"
	//address = "Sleman"
	//if address == "Turi" {
	//	fmt.Printf("Address : %s %d Tahun\n", address, age)
	//} else {
	//	fmt.Printf("Address is not Turi, but %s %d Tahun\n", address, age)
	//}
	//fmt.Printf("My Name : %s\n", name)
	//
	//// Seed the random number generator with the current time
	////rand.Seed(time.Now().UnixNano())
	//
	//// Define the minimum and maximum values
	//min := 1.0
	//max := 130.0
	//
	//// Generate a random floating-point number within the specified range
	//var point = rand.Float64()*(max-min) + min
	//
	//if percent := point; percent >= 100.0 {
	//	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	//} else if percent >= 70.0 {
	//	fmt.Printf("%.1f%s good\n", percent, "%")
	//} else {
	//	fmt.Printf("%.1f%s not bad | point : %f\n", percent, "%", point)
	//}
	//
	//var i = 0
	//
	//for i < 5 {
	//	if i == 3 {
	//		break
	//	}
	//	fmt.Println("Angka", i)
	//	i++
	//}
	//
	//var names = [...]string{"rahma", "vani", "aan"}
	//
	//for i, name := range names {
	//	fmt.Printf(name+" %d \n ", i)
	//}
	//
	//for i, _ := range names {
	//	fmt.Printf(names[i]+" %d \n", i)
	//}
	//
	//fmt.Println(len(names[0:2]))
	//
	//for _, name := range append(names[0:2], "olesky") {
	//	println(name)
	//}

	//var names = []string{"aan", "vani"}
	//
	//for _, name := range names {
	//	println(name)
	//}
	//
	//oriArray := []string{"Rahma", "Purnama"}
	//
	//var copyArray []string
	//
	//oriArray[0] = "Turi"
	//
	//copy(copyArray, oriArray)
	//
	//for _, name := range copyArray {
	//	println(name)
	//}
	//
	//animals := map[string]any{
	//	"darat": 1,
	//	"laut":  "Salmon",
	//	"udara": 3,
	//}
	//
	//var _, isExist = animals["udara"]
	//
	//if isExist {
	//	delete(animals, "udara")
	//}
	//
	//for key, value := range animals {
	//	println("Key : "+key+" | Value : ", fmt.Sprintf("%v", value))
	//}

	multiArray := []map[string]any{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	PrintArrayMulti("======================", multiArray)

	findWord := func(words string, cb func(string) bool, cb1 func(string) string) string {
		if cb(words) {
			return cb1(words) + " Found!"
		}

		return cb1(words)
	}

	println(findWord("Rahma Andita Purnama", func(s string) bool {
		return strings.Contains(strings.ToLower(s), "rahma")
	}, func(s string) string {
		return strings.ToUpper(s)
	}))

	var number1 = 1

	var number2 = &number1

	number1 = 10

	println(*number2)

	type Class struct {
		name    string
		lecture string
	}
	type Student struct {
		name    string
		age     int
		address string
		class   Class
	}

	//var user1 Student
	//
	//user1.name = "Rahma"
	//user1.age = 17
	//user1.address = "Turi"
	//println(user1.address, user1.name, user1.age)
	class1 := Class{name: "SI-07", lecture: "Doni"}
	user2 := Student{name: "Rahmaap", age: 19, address: "Turi", class: class1}

	println(user2.age, user2.class.name)
}

func PrintArrayMulti(message string, arr []map[string]any) {
	println(message)
	for _, value := range arr {
		for key, v := range value {
			println("Key : "+key+" | Value :", fmt.Sprintf("%v", v))
		}
	}

	println(persegi(10, 20))
	calculate := calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Printf("%.2f\n\n", calculate)
}

func persegi(panjang, lebar int) (string, int) {
	return "Luas Persegi adalah : ", panjang * lebar
}

func calculate(numbers ...int) (avg float64) {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	avg = float64(total) / float64(len(numbers))

	return
}
