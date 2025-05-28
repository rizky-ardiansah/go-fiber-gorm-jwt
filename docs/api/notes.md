# Dokumentasi Notes API

Dokumentasi ini menjelaskan endpoint API yang berkaitan dengan Notes. Semua endpoint di bawah ini memerlukan autentikasi JWT melalui HTTP-only cookies.

## Membuat Catatan Baru

Digunakan untuk membuat catatan baru milik pengguna yang terautentikasi.

- **URL**: `/api/notes`
- **Method**: `POST`
- **Protected**: Ya

### Headers Permintaan

| Header   | Deskripsi                                                  | Contoh    |
| -------- | ---------------------------------------------------------- | --------- |
| `Cookie` | Cookie yang berisi token JWT (`jwt=your_jwt_token_string`) | `jwt=...` |

**Catatan**: Autentikasi menggunakan HTTP-only cookies, bukan Authorization header.

### Body Permintaan (JSON)

```json
{
  "title": "Judul Catatan Baru",
  "content": "Isi dari catatan baru."
}
```

### Respons Sukses

- **Kode**: `201 CREATED`
- **Konten**: Objek JSON yang berisi status, pesan, dan data catatan yang baru dibuat.

  ```json
  {
    "status": "success",
    "message": "Note created successfully",
    "data": {
      "ID": 1,
      "CreatedAt": "2025-05-28T10:00:00Z",
      "UpdatedAt": "2025-05-28T10:00:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan Baru",
      "content": "Isi dari catatan baru.",
      "user_id": 123
    }
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika data tidak valid)
  ```json
  {
    "error": "Cannot parse JSON"
  }
  ```
- **Kode**: `401 Unauthorized` (jika pengguna tidak terautentikasi)
  ```json
  {
    "error": "Unauthorized: User ID not found in context"
  }
  ```
- **Kode**: `500 Internal Server Error` (jika terjadi kesalahan server)
  ```json
  {
    "error": "Could not create note"
  }
  ```

## Mendapatkan Semua Catatan Pengguna

Digunakan untuk mengambil semua catatan milik pengguna yang terautentikasi.

- **URL**: `/api/notes`
- **Method**: `GET`
- **Protected**: Ya

### Headers Permintaan

| Header   | Deskripsi                                                            | Contoh    |
| -------- | -------------------------------------------------------------------- | --------- |
| `Cookie` | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`) | `jwt=...` |

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Array objek JSON yang berisi semua catatan pengguna (response langsung tanpa wrapper).

  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2025-05-28T10:00:00Z",
      "UpdatedAt": "2025-05-28T10:00:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan Pertama",
      "content": "Isi catatan pertama.",
      "user_id": 123
    },
    {
      "ID": 2,
      "CreatedAt": "2025-05-28T10:05:00Z",
      "UpdatedAt": "2025-05-28T10:05:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan Kedua",
      "content": "Isi catatan kedua.",
      "user_id": 123
    }
  ]
  ```

**Catatan**: Jika tidak ada catatan, akan mengembalikan array kosong `[]`.

### Respons Error

- **Kode**: `401 Unauthorized`
  ```json
  {
    "error": "Unauthorized: User ID not found in context"
  }
  ```
- **Kode**: `500 Internal Server Error`
  ```json
  {
    "error": "Could not retrieve notes"
  }
  ```

## Mendapatkan Detail Catatan

Digunakan untuk mengambil detail catatan spesifik berdasarkan ID.

- **URL**: `/api/notes/:id`
- **Method**: `GET`
- **Protected**: Ya

### Headers Permintaan

| Header   | Deskripsi                                                            | Contoh    |
| -------- | -------------------------------------------------------------------- | --------- |
| `Cookie` | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`) | `jwt=...` |

### Parameter URL

| Parameter | Deskripsi   | Contoh |
| --------- | ----------- | ------ |
| `id`      | ID catatan. | `1`    |

### Respons Sukses

- **Kode**: `200 OK`
- **Konten**: Objek JSON yang berisi status, pesan, dan data catatan.
  ```json
  {
    "status": "success",
    "message": "Note retrieved successfully",
    "data": {
      "ID": 1,
      "CreatedAt": "2025-05-28T10:00:00Z",
      "UpdatedAt": "2025-05-28T10:00:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan",
      "content": "Isi catatan.",
      "user_id": 123
    }
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika ID tidak valid)
  ```json
  {
    "error": "Invalid note ID"
  }
  ```
- **Kode**: `401 Unauthorized`
  ```json
  {
    "error": "Unauthorized: User ID not found in context"
  }
  ```
- **Kode**: `404 Not Found` (jika catatan tidak ditemukan atau bukan milik pengguna)
  ```json
  {
    "error": "Note not found"
  }
  ```
- **Kode**: `500 Internal Server Error`
  ```json
  {
    "error": "Internal Server Error: User ID is of an invalid type"
  }
  ```

## Memperbarui Catatan

Digunakan untuk memperbarui catatan yang sudah ada.

- **URL**: `/api/notes/:id`
- **Method**: `PUT`
- **Protected**: Ya

### Headers Permintaan

| Header   | Deskripsi                                                            | Contoh    |
| -------- | -------------------------------------------------------------------- | --------- |
| `Cookie` | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`) | `jwt=...` |

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
- **Konten**: Objek JSON yang berisi status, pesan, dan data catatan yang telah diperbarui.

  ```json
  {
    "status": "success",
    "message": "Note updated successfully",
    "data": {
      "ID": 1,
      "CreatedAt": "2025-05-28T10:00:00Z",
      "UpdatedAt": "2025-05-28T10:15:00Z",
      "DeletedAt": null,
      "title": "Judul Catatan Diperbarui",
      "content": "Isi dari catatan yang telah diperbarui.",
      "user_id": 123
    }
  }
  ```

### Respons Error

- **Kode**: `400 Bad Request` (jika ID atau data tidak valid)
  ```json
  {
    "error": "Invalid note ID"
  }
  ```
  atau
  ```json
  {
    "error": "Cannot parse JSON"
  }
  ```
- **Kode**: `401 Unauthorized`
  ```json
  {
    "error": "Unauthorized: User ID not found in context"
  }
  ```
- **Kode**: `404 Not Found` (jika catatan tidak ditemukan atau bukan milik pengguna)
  ```json
  {
    "error": "Note not found"
  }
  ```
- **Kode**: `500 Internal Server Error`
  ```json
  {
    "error": "Could not update note"
  }
  ```

## Menghapus Catatan

Digunakan untuk menghapus catatan.

- **URL**: `/api/notes/:id`
- **Method**: `DELETE`
- **Protected**: Ya

### Headers Permintaan

| Header   | Deskripsi                                                            | Contoh    |
| -------- | -------------------------------------------------------------------- | --------- |
| `Cookie` | Cookie yang berisi token JWT (misalnya, `jwt=your_jwt_token_string`) | `jwt=...` |

### Parameter URL

| Parameter | Deskripsi   | Contoh |
| --------- | ----------- | ------ |
| `id`      | ID catatan. | `1`    |

### Respons Sukses

- **Kode**: `204 No Content`

**Catatan**: Response tidak memiliki body content saat berhasil menghapus.

### Respons Error

- **Kode**: `400 Bad Request` (jika ID tidak valid)
  ```json
  {
    "error": "Invalid note ID"
  }
  ```
- **Kode**: `401 Unauthorized`
  ```json
  {
    "error": "Unauthorized: User ID not found in context"
  }
  ```
- **Kode**: `404 Not Found` (jika catatan tidak ditemukan atau bukan milik pengguna)
  ```json
  {
    "error": "Note not found"
  }
  ```
- **Kode**: `500 Internal Server Error`
  ```json
  {
    "error": "Could not delete note"
  }
  ```
