# Dokumentasi API User

Dokumentasi ini menjelaskan endpoint API yang berkaitan dengan pengguna. Endpoint ini menggunakan HTTP-only cookies untuk autentikasi.

## Mendapatkan Profil Pengguna Saya

Digunakan untuk mengambil detail profil pengguna yang saat ini terautentikasi.

- **URL**: `/api/users/me`
- **Method**: `GET`
- **Protected**: Ya (Membutuhkan token JWT)

### Headers Permintaan

| Header   | Deskripsi                                                     | Contoh    |
| -------- | ------------------------------------------------------------- | --------- |
| `Cookie` | Cookie yang berisi token JWT (`jwt=your_jwt_token_string`)   | `jwt=...` |

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi status, pesan, dan data pengguna.

  ```json
  {
    "status": "success",
    "message": "Profile fetched successfully",
    "data": {
      "ID": 1,
      "CreatedAt": "2025-05-28T10:00:00Z",
      "UpdatedAt": "2025-05-28T10:00:00Z",
      "DeletedAt": null,
      "name": "Nama Pengguna",
      "email": "user@example.com",
      "password": ""
    }
  }
  ```

  **Catatan**: Field `password` akan selalu dikosongkan demi keamanan.

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
- **Kode**: `404 Not Found` (jika user tidak ditemukan)
  ```json
  {
    "status": "error",
    "message": "User not found"
  }
  ```

  _(Pesan error spesifik mungkin berbeda tergantung implementasi middleware `Protected`)_
