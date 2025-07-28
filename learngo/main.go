// main.go : for compile
package main

import (
	"fmt"
	"strings"

	"github.com/t4e1/go-prac1/learngo/something"
)

// Functions
// argument, return type지정해줘야함
func multiply(a int, b int) int {
	return a * b
}

// go에서는 return 값을 여러개 가져갈 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// unlimited argument 받는 방법 ...type
func repeatMe(words ...string) {
}

// naked return : return 에 variable을 명시하지 않아도 되게함. 함수의 리턴값을 정의할 때 미리 리턴 variable을 정해놨기 때문
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.") // return 한 뒤에 function 이 끝났을 때 작업을 실행시킬 수 있음
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// Loop
// Go에서의 반복문은 for 문만 사용
// range는 index를 준다
func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// // }
	// for i:= 0; i< len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	total := 0
	for _, number := range numbers { // _는 리턴을 무시하고 싶을 때 사용함
		total += number
	}
	return total
}

// If Else
func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge < 18 { // if 사용시 variable 을 생성할 수 있음 ( koreanAge)
		return false
	}
	return true
}

// Switch
func canIDrink2(age int) bool {
	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// case age > 50:
	// 	return false
	// }

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// Struct
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {

	//Packages and Import, Export
	fmt.Println("Hello world!")
	something.SayHello()

	// //Variables
	// const name string = "tw" // const = 상수
	// var name2 string = "tw"  // var = 변수
	// name3 := "tw"            // shorthand = 자동으로 첫 값으로 타입을 정해줌. func 안에서만 동작하고 밖에서는 타입 선언 필요
	// name2 = "wt"
	// fmt.Println(name2)

	//Functions
	fmt.Println(multiply(2, 2))
	totlaLength, upperName := lenAndUpper("tw")
	fmt.Println(totlaLength, upperName)
	repeatMe("a", "b", "c", "d", "e")

	//Functions2
	totalLength2, up := lenAndUpper2("wt")
	fmt.Println(totalLength2, up)

	//Loop
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)

	//If & Else
	fmt.Println(canIDrink(16))

	//Switch
	fmt.Println(canIDrink2(16))

	//Pointer
	// Lowlevel Programming
	// 대규모 시스템에서 계속 copy를 만들지 않고 하나의 Object를 활용할 수 있음.
	// & :  memory address
	// * : see through
	a := 2
	// b := a
	b := &a // &를 사용하면 변수의 메모리 주소를 불러올 수 있음
	a = 5
	*b = 20 // *는 a의 주소값을 열기 때문에 a = 20 이 된다
	fmt.Println(a, b)
	fmt.Println(&a, b, &b)
	fmt.Println(a, *b) // *를 사용하면 b에 저장된 메모리주소(a의 메모리주소)에 저장된 값을 볼 수 있음

	// Arrays
	names := [5]string{"tw", "wt", "dal"}
	names[3] = "dil"

	names2 := []string{"tw", "wt", "dal"} // slice : 길이를 지정하지 않는 array
	// names2[3] = "dil" // 에러 발생함 (array 길이를 넘는다는 에러) 값을 추가하고 싶을떈 append 로 추가해줘야함.
	names2 = append(names2, "dil")
	fmt.Println(names, names2)

	// Maps
	// Key - Value 가 있음
	// [] 안에 Key 타입이, 뒤에는 Value 타입이 들어감
	tw := map[string]string{"name": "tw", "age": "32"}
	for key, value := range tw { // map 도 for range 를 활용할 수 있음
		fmt.Println(key, value)
	}
	fmt.Println(tw)

	// Struct
	// Object와 비슷하면서 map보다 좀 더 유연함
	favFood := []string{"ramen", "kimchi"}
	// tw3 := person{"tw", 11, favFood}
	tw3 := person{name: "tw", age: 11, favFood: favFood} // 각 요소들을 잘 보여주기 위해서 name , age, favFood 명시해주기
	fmt.Println(tw3, tw3.name)
}
