# Dokumentasi API Autentikasi

Dokumentasi ini menjelaskan endpoint API yang berkaitan dengan autentikasi pengguna. Semua endpoint autentikasi menggunakan HTTP-only cookies untuk menyimpan JWT token demi keamanan yang lebih baik.

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
- **Konten**: Objek JSON yang berisi status, pesan, dan data pengguna. JWT token juga disimpan dalam HTTP-only cookie.

  ```json
  {
    "status": "success",
    "message": "user registered successfully",
    "data": {
      "ID": 1,
      "CreatedAt": "2025-05-28T10:00:00Z",
      "UpdatedAt": "2025-05-28T10:00:00Z",
      "DeletedAt": null,
      "name": "Nama Pengguna",
      "email": "user@example.com",
      "password": "$2a$10$..." // password hash
    }
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika data tidak valid atau field required kosong)
  ```json
  {
    "status": "error",
    "message": "name, email, and password are required"
  }
  ```
- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan server)
  ```json
  {
    "status": "error",
    "message": "failed to create user",
    "data": "error details"
  }
  ```

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
- **Konten**: Objek JSON yang berisi status, pesan, dan data token. Token juga akan disimpan dalam HTTP-only cookie.

  ```json
  {
    "status": "success",
    "message": "Login successful",
    "data": {
      "token": "your_jwt_token_string"
    }
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika email atau password tidak disediakan)
  ```json
  {
    "status": "error",
    "message": "Email and Password are required"
  }
  ```
- **Kode**: `401 Unauthorized` (jika kredensial salah)
  ```json
  {
    "status": "error",
    "message": "Invalid email or password"
  }
  ```
- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan server)
  ```json
  {
    "status": "error",
    "message": "Could not generate token",
    "data": "error details"
  }
  ```

## Logout Pengguna

Digunakan untuk menghapus cookie JWT pengguna, secara efektif melakukan logout.

- **URL**: `/api/auth/logout`
- **Method**: `POST`
- **Protected**: Tidak (Namun, idealnya hanya bisa diakses oleh pengguna yang sudah login)

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi status dan pesan konfirmasi. HTTP-only cookie JWT akan dihapus.

  ```json
  {
    "status": "success",
    "message": "Logged out successfully"
  }
  ```

### Respons Error

- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan saat menghapus cookie)
