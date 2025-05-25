# Dokumentasi API Proyek Go Fiber GORM JWT

Selamat datang di dokumentasi API untuk proyek ini. API ini dibangun menggunakan Go, Fiber, GORM, dan menggunakan JWT untuk autentikasi.

## Modul API

Berikut adalah modul-modul API yang tersedia:

-   [Autentikasi (Auth)](auth.md) - Endpoint untuk registrasi, login, dan logout pengguna.
-   [Pengguna (Users)](users.md) - Endpoint untuk manajemen data pengguna.
-   [Catatan (Notes)](notes.md) - Endpoint untuk operasi CRUD pada catatan pengguna.

## Autentikasi

Sebagian besar endpoint API memerlukan autentikasi menggunakan JSON Web Tokens (JWT).
Setelah berhasil login, token JWT akan dikirimkan kembali kepada klien dan juga disimpan dalam HTTP-only cookie bernama `jwt`.
Untuk permintaan selanjutnya ke endpoint yang terproteksi, klien harus menyertakan token ini:
-   Secara otomatis melalui cookie yang dikirimkan oleh browser.
-   Atau, secara manual dengan menyertakan header `Authorization: Bearer <your_jwt_token>`. Middleware `Protected` akan memeriksa cookie terlebih dahulu, kemudian header `Authorization`.

## Struktur Respons Umum

### Sukses
Respons sukses umumnya akan memiliki kode status HTTP `200 OK` atau `201 Created` dan berisi data yang diminta dalam format JSON.

### Error
Respons error akan memiliki kode status HTTP yang sesuai (misalnya, `400 Bad Request`, `401 Unauthorized`, `404 Not Found`, `500 Internal Server Error`) dan biasanya berisi objek JSON dengan detail error:

```json
{
    "error": "Deskripsi singkat mengenai error"
}
```
atau
```json
{
    "status": "error",
    "message": "Deskripsi error yang lebih detail"
}
```
Pesan error spesifik dapat bervariasi tergantung pada endpoint dan jenis kesalahan.
