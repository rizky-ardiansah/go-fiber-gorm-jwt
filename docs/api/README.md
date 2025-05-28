# Dokumentasi API Proyek Go Fiber GORM JWT

Selamat datang di dokumentasi API untuk proyek ini. API ini dibangun menggunakan Go, Fiber, GORM, dan menggunakan JWT dengan HTTP-only cookies untuk autentikasi yang aman.

## Base URL

- **Development**: `http://localhost:8080`

## Modul API

Berikut adalah modul-modul API yang tersedia:

- [Autentikasi (Auth)](auth.md) - Endpoint untuk registrasi, login, dan logout pengguna menggunakan HTTP-only cookies.
- [Pengguna (Users)](users.md) - Endpoint untuk manajemen data pengguna yang terautentikasi.
- [Catatan (Notes)](notes.md) - Endpoint untuk operasi CRUD pada catatan pengguna.

## Autentikasi

API ini menggunakan **HTTP-only cookies** untuk autentikasi JWT, yang memberikan keamanan yang lebih baik dibanding token di localStorage:

### Keunggulan HTTP-only Cookies:

- ✅ **Anti-XSS**: Cookie tidak dapat diakses melalui JavaScript
- ✅ **Anti-CSRF**: Menggunakan SameSite policy
- ✅ **Secure**: Cookie hanya dikirim melalui HTTPS di production
- ✅ **Automatic**: Browser mengelola cookie secara otomatis

### Konfigurasi Cookie:

```javascript
{
  Name: "jwt",
  Path: "/",
  MaxAge: 86400, // 1 hari
  Secure: true,  // true di production
  HTTPOnly: true,
  SameSite: "Lax"
}
```

### Cara Kerja:

1. **Login**: POST ke `/api/auth/login` dengan credentials
2. **Cookie Setting**: Server menyimpan JWT dalam HTTP-only cookie
3. **Request Berikutnya**: Browser otomatis menyertakan cookie
4. **Logout**: POST ke `/api/auth/logout` untuk menghapus cookie

## Struktur Respons Umum

### Sukses

Respons sukses memiliki format yang konsisten:

**Auth & Users Endpoints:**

```json
{
  "status": "success",
  "message": "Descriptive success message",
  "data": {
    /* actual data */
  }
}
```

**Notes Endpoints:**

- **GET /api/notes**: Mengembalikan array langsung `[...]`
- **POST/PUT /api/notes**: Menggunakan format wrapper dengan status
- **DELETE /api/notes/:id**: Mengembalikan status `204 No Content`

### Error

Respons error konsisten menggunakan format:

```json
{
  "error": "Deskripsi error yang jelas"
}
```

atau untuk beberapa endpoint auth:

```json
{
  "status": "error",
  "message": "Deskripsi error yang detail",
  "data": "additional error info (optional)"
}
```

### Status Codes

- `200 OK`: Request berhasil
- `201 Created`: Resource berhasil dibuat
- `204 No Content`: Request berhasil tanpa content (biasanya DELETE)
- `400 Bad Request`: Request tidak valid
- `401 Unauthorized`: Autentikasi gagal atau tidak ada
- `404 Not Found`: Resource tidak ditemukan
- `500 Internal Server Error`: Error di server
