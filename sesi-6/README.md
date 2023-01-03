# SESI 06
- Kode Peserta  : 149368582100-676
- Nama          : Vini Jumatul Fitri 

## Summary sesi 6 : 
1. *Web Server* -> untuk membuat aplikasi berbasis web pada bahasa go telah menyediakan package bernama net/http.

### Cara running program
* Link Postman : https://www.getpostman.com/collections/5c49fed3e57390049061

1. Buka terminal dan masuk ke folder sesi-06
```
cd sesi-06/web-server
```
2. Lakukan perintah running untuk memulai server
```
go run api.go
```
3. Lakukan hit api pada postman / website
```
METHOD : GET
URL : localhost:8080/employees

METHOD : POST
URL : localhost:8080/employee
```


2. *Gin Framework* -> sebuah framework bahasa Go yang digunakan untuk keperluan http routing.


### Cara running program
 * Link Postman : https://www.getpostman.com/collections/a40fe6890005285c4042

1. Buka terminal dan masuk ke folder sesi-06
```
cd sesi-06/gin-fw
```
2. Lakukan perintah running untuk memulai server
```
go run main.go
```
3. Lakukan hit api pada postman / website
```
METHOD : POST
URL : localhost:8080/cars

METHOD : GET
URL : localhost:8080/cars/:carID

METHOD : PUT
URL : localhost:8080/cars/:carID

METHOD : DELETE
URL : localhost:8080/cars/:carID
```