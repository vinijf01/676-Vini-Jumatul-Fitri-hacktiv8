package main

import "fmt"

func main() {
	/*0. Slice*/

	var fruits = []string{"apple", "banana", "mango", "pear"}
	_ = fruits

	/*1. Slice(use make function)*/
	var buahs = make([]string, 3)
	_ = buahs

	fmt.Printf("%#v\n", buahs)

	/*2. Slice (append function)*/
	//meng entri kan value ke index yang kosong
	buahs[0] = "one"
	buahs[1] = "two"
	buahs[2] = "three"
	fmt.Printf("%#v\n", buahs)

	//menambahkan nilai menggunakan fungsi append
	buahs = append(buahs, "apel", "pisang", "mangga")
	fmt.Printf("%#v\n", buahs)

	/*3. Slice (append function with ellipsis)*/
	var fruits1 = []string{"Apple", "Banana", "Mango"}
	var fruits2 = []string{"Durian", "pineapple", "startfruit"}

	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("%#v\n", fruits1)

	/*4. Slice (copy function)*/

	var buah1 = []string{"apple", "banana", "mango"}
	var buah2 = []string{"durian", "pineapple", "startfruit"}

	nn := copy(buah1, buah2)
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println("copy element =>", nn)

	/*5. Slice(slicing)*/

	var fruit1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruit2 = fruit1[1:2]
	fmt.Printf("%#v\n", fruit2)

	var fruit3 = fruit1[0:]
	fmt.Printf("%#v\n", fruit3)

	var fruit4 = fruit1[:3]
	fmt.Printf("%#v\n", fruit4)

	var fruit5 = fruit1[:]
	fmt.Printf("%#v\n", fruit5)

	/*6. Slice (Combining slicing and append)*/

	var fruit6 = append(fruit1[:3], "rambutan")
	fmt.Printf("%#v\n", fruit6)

	/*7. Slice(Backing array)*/

	var fruit7 = fruit1[2:4]

	fruit7[0] = "cery"
	fmt.Println("fruit1 =>", fruit1)
	fmt.Println("fruit7 =>", fruit7)

	/*8. Slice (Cap function)*/
	var fruitSatu = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("FruiteSatu cap : ", cap(fruitSatu))
	fmt.Println("FruitSatu leng : ", len(fruitSatu))

	var fruitdua = fruitSatu[0:3]

	fmt.Println("FruiteDua cap : ", cap(fruitdua))
	fmt.Println("FruitDua leng : ", len(fruitdua))

	var fruitTiga = fruitSatu[1:]
	fmt.Println("FruiteTiga cap : ", cap(fruitTiga)) //fruit[x:y] cap bisa berubah karena nilai x yang > 0, membuat elemen ke-x slice yang diambil menjadi elemen ke 0 slice baru
	fmt.Println("FruitTiga leng : ", len(fruitTiga))

	/*9. Slice (Creating a new backing array)*/
	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "nissan"
	fmt.Println("cars: ", cars)
	fmt.Println("newCars: ", newCars)
}
