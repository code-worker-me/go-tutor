#### Usage
Setiap kali menambah tag annotation swagger pada handler wajib menjalankan perintah berikut untuk memperbarui API docs.

```bash
swag init -g cmd/server/main.go
```

Untuk menjalankan app, jelankan command berikut:
```bash
go run cmd/server/main.go
```

buka link berikut untuk melihat api docs.
```http
http://localhost:8080/swagger/index.html
```