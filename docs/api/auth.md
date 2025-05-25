# Dokumentasi API Autentikasi

Dokumentasi ini menjelaskan endpoint API yang berkaitan dengan autentikasi pengguna.

## Registrasi Pengguna Baru

Digunakan untuk mendaftarkan pengguna baru ke dalam sistem.

- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Protected**: Tidak

### Body Permintaan (JSON)

```json
{
  "name": "Nama Pengguna",
  "email": "user@example.com",
  "password": "passwordkuat"
}
```

### Respons Sukses

- **Kode**: `201 CREATED`
- **Konten**: Objek JSON yang berisi detail pengguna yang baru dibuat (tanpa password).

  ```json
  {
    "id": 1,
    "name": "Nama Pengguna",
    "email": "user@example.com"
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika data tidak valid atau email sudah terdaftar)
- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan server)

## Login Pengguna

Digunakan untuk mengautentikasi pengguna yang sudah terdaftar dan mendapatkan token JWT.

- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Protected**: Tidak

### Body Permintaan (JSON)

```json
{
  "email": "user@example.com",
  "password": "passwordkuat"
}
```

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi token JWT. Token juga akan disimpan dalam HTTP-only cookie.

  ```json
  {
    "token": "your_jwt_token_string"
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika email atau password tidak valid)
- **Kode**: `401 Unauthorized` (jika kredensial salah)
- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan server)

## Logout Pengguna

Digunakan untuk menghapus cookie JWT pengguna, secara efektif melakukan logout.

- **URL**: `/api/auth/logout`
- **Method**: `POST`
- **Protected**: Tidak (Namun, idealnya hanya bisa diakses oleh pengguna yang sudah login)

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Pesan sukses.

  ```json
  {
    "message": "Logout successful"
  }
  ```

### Respons Error

- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan saat menghapus cookie)
