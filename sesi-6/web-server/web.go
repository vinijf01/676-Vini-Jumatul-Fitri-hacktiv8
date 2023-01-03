package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet) //routing

	http.ListenAndServe(PORT, nil) //untuk menjalankan server aplikasi
}

/*
- http.ResponWriter : sebuah interface dengan berbagai method yang digunakan
                      untuk mengirim response balik kpd client yg mengirim request
- http.Request : sebuah struct yang digunakan untuk mendapat berbagai macam data request
*/

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "hello World"
	fmt.Fprint(w, msg)
}
