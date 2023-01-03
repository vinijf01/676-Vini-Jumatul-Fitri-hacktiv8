package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Grade int
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue1 = reflect.ValueOf(number)

	fmt.Println("tipe variable :", reflectValue1.Type())

	fmt.Println("Nilai varabel :", reflectValue1.Interface()) // jika nilai hanya diperlukan untuk ditampilkan use method interface()

	if reflectValue1.Kind() == reflect.Int {
		fmt.Println("Nilai varabel :", reflectValue1.Int())

	}

	/// identifying Method Information
	var s1 = &Student{Name: "Vini Jumatul", Grade: 80}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Jumatul"),
	})

	fmt.Println("nama : ", s1.Name)
}
