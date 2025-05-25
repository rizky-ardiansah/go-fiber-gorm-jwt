# Dokumentasi Notes API

Dokumentasi ini menjelaskan endpoint API yang berkaitan dengan Notes. Semua endpoint di bawah ini memerlukan autentikasi JWT.

## Membuat Catatan Baru

Digunakan untuk membuat catatan baru milik pengguna yang terautentikasi.

- **URL**: `/api/notes`
- **Method**: `POST`
- **Protected**: Ya

### Headers Permintaan

| Header          | Deskripsi                                                                     | Contoh                    |
| --------------- | ----------------------------------------------------------------------------- | ------------------------- |
| `Cookie`        | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`)          | `jwt=...`                 |
| `Authorization` | Opsional: Token JWT untuk autentikasi pengguna jika tidak menggunakan cookie. | `Bearer <your_jwt_token>` |

### Body Permintaan (JSON)

```json
{
  "title": "Judul Catatan Baru",
  "content": "Isi dari catatan baru."
}
```

### Respons Sukses

- **Kode**: `201 CREATED`
- **Konten**: Objek JSON yang berisi detail catatan yang baru dibuat.

  ```json
  {
    "ID": 1,
    "CreatedAt": "2025-05-25T10:00:00Z",
    "UpdatedAt": "2025-05-25T10:00:00Z",
    "DeletedAt": null,
    "title": "Judul Catatan Baru",
    "content": "Isi dari catatan baru.",
    "user_id": 123
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika data tidak valid)
- **Kode**: `401 Unauthorized` (jika pengguna tidak terautentikasi)
- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan server)

## Mendapatkan Semua Catatan Pengguna

Digunakan untuk mengambil semua catatan milik pengguna yang terautentikasi.

- **URL**: `/notes`
- **Method**: `GET`
- **Protected**: Ya

### Headers Permintaan

| Header          | Deskripsi                                                                     | Contoh                    |
| --------------- | ----------------------------------------------------------------------------- | ------------------------- |
| `Cookie`        | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`)          | `jwt=...`                 |
| `Authorization` | Opsional: Token JWT untuk autentikasi pengguna jika tidak menggunakan cookie. | `Bearer <your_jwt_token>` |

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Array objek JSON yang berisi semua catatan pengguna.

  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2025-05-25T10:00:00Z",
      "UpdatedAt": "2025-05-25T10:00:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan Pertama",
      "content": "Isi catatan pertama.",
      "user_id": 123
    },
    {
      "ID": 2,
      "CreatedAt": "2025-05-25T10:05:00Z",
      "UpdatedAt": "2025-05-25T10:05:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan Kedua",
      "content": "Isi catatan kedua.",
      "user_id": 123
    }
  ]
  ```

### Respons Error

- **Kode**: `401 Unauthorized`
- **Kode**: `500 Internal Server Error`

## Mendapatkan Detail Catatan

Digunakan untuk mengambil detail catatan spesifik berdasarkan ID.

- **URL**: `/api/notes/:id`
- **Method**: `GET`
- **Protected**: Ya

### Headers Permintaan

| Header          | Deskripsi                                                                     | Contoh                    |
| --------------- | ----------------------------------------------------------------------------- | ------------------------- |
| `Cookie`        | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`)          | `jwt=...`                 |
| `Authorization` | Opsional: Token JWT untuk autentikasi pengguna jika tidak menggunakan cookie. | `Bearer <your_jwt_token>` |

### Parameter URL

| Parameter | Deskripsi   | Contoh |
| --------- | ----------- | ------ |
| `id`      | ID catatan. | `1`    |

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi detail catatan.

  ```json
  {
    "ID": 1,
    "CreatedAt": "2025-05-25T10:00:00Z",
    "UpdatedAt": "2025-05-25T10:00:00Z",
    "DeletedAt": null,
    "title": "Judul Catatan",
    "content": "Isi catatan.",
    "user_id": 123
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika ID tidak valid)
- **Kode**: `401 Unauthorized`
- **Kode**: `404 Not Found` (jika catatan tidak ditemukan atau bukan milik pengguna)
- **Kode**: `500 Internal Server Error`

## Memperbarui Catatan

Digunakan untuk memperbarui catatan yang sudah ada.

- **URL**: `/api/notes/:id`
- **Method**: `PUT`
- **Protected**: Ya

### Headers Permintaan

| Header          | Deskripsi                                                                     | Contoh                    |
| --------------- | ----------------------------------------------------------------------------- | ------------------------- |
| `Cookie`        | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`)          | `jwt=...`                 |
| `Authorization` | Opsional: Token JWT untuk autentikasi pengguna jika tidak menggunakan cookie. | `Bearer <your_jwt_token>` |

### Parameter URL

| Parameter | Deskripsi   | Contoh |
| --------- | ----------- | ------ |
| `id`      | ID catatan. | `1`    |

### Body Permintaan (JSON)

```json
{
  "title": "Judul Catatan Diperbarui",
  "content": "Isi dari catatan yang telah diperbarui."
}
```

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi detail catatan yang telah diperbarui.

  ```json
  {
    "ID": 1,
    "CreatedAt": "2025-05-25T10:00:00Z",
    "UpdatedAt": "2025-05-25T10:15:00Z",
    "DeletedAt": null,
    "title": "Judul Catatan Diperbarui",
    "content": "Isi dari catatan yang telah diperbarui.",
    "user_id": 123
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika ID atau data tidak valid)
- **Kode**: `401 Unauthorized`
- **Kode**: `404 Not Found` (jika catatan tidak ditemukan atau bukan milik pengguna)
- **Kode**: `500 Internal Server Error`

## Menghapus Catatan

Digunakan untuk menghapus catatan.

- **URL**: `/api/notes/:id`
- **Method**: `DELETE`
- **Protected**: Ya

### Headers Permintaan

| Header          | Deskripsi                                                                     | Contoh                    |
| --------------- | ----------------------------------------------------------------------------- | ------------------------- |
| `Cookie`        | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`)          | `jwt=...`                 |
| `Authorization` | Opsional: Token JWT untuk autentikasi pengguna jika tidak menggunakan cookie. | `Bearer <your_jwt_token>` |

### Parameter URL

| Parameter | Deskripsi   | Contoh |
| --------- | ----------- | ------ |
| `id`      | ID catatan. | `1`    |

### Respons Sukses

- **Kode**: `204 No Content`

### Respons Error

- **Kode**: `400 Bad Request` (jika ID tidak valid)
- **Kode**: `401 Unauthorized`
- **Kode**: `404 Not Found` (jika catatan tidak ditemukan atau bukan milik pengguna)
- **Kode**: `500 Internal Server Error`
