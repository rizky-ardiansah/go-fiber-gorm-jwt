# Dokumentasi API User

Dokumentasi ini menjelaskan endpoint API yang berkaitan dengan pengguna.

## Mendapatkan Profil Pengguna Saya

Digunakan untuk mengambil detail profil pengguna yang saat ini terautentikasi.

- **URL**: `/api/users/me`
- **Method**: `GET`
- **Protected**: Ya (Membutuhkan token JWT)

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi detail pengguna.

  ```json
  {
    "id": 1,
    "name": "Nama Pengguna",
    "email": "user@example.com",
    "created_at": "2025-05-25T10:00:00Z",
    "updated_at": "2025-05-25T10:00:00Z"
  }
  ```

  _(Catatan: Struktur data pengguna mungkin berbeda tergantung pada model `User` Anda)_

### Respons Error

- **Kode**: `401 Unauthorized`
  - **Deskripsi**: Terjadi jika token JWT tidak valid, kedaluwarsa, atau tidak disertakan.
  - **Konten**:
    ```json
    {
      "status": "error",
      "message": "Invalid or expired token"
    }
    ```
    _(Pesan error spesifik mungkin berbeda tergantung implementasi middleware `Protected`)_
