# Go Fiber GORM JWT - Notes API

Simple REST API untuk aplikasi catatan (notes) yang dibangun menggunakan Go, Fiber, GORM, dan JWT authentication.

---

## Fitur

- ✅ Autentikasi JWT (Register, Login, Logout)
- ✅ CRUD Catatan (Create, Read, Update, Delete)
- ✅ Middleware proteksi untuk endpoint yang memerlukan autentikasi
- ✅ Relasi database antara User dan Notes
- ✅ Validasi input dan error handling

---

## Tech Stack

- **Framework**: [Fiber](https://gofiber.io/)
- **Database ORM**: [GORM](https://gorm.io/)
- **Authentication**: JWT (JSON Web Tokens)
- **Database**: PostgreSQL
- **Password Hashing**: bcrypt

---

## Instalasi

1. Clone repository:

```bash
git clone <repository-url>
cd go-fiber-gorm-jwt
```

2. Install dependencies:

```bash
go mod tidy
```

3. Setup environment variables:

```bash
cp .env.example .env
```

Edit file `.env` dan sesuaikan dengan konfigurasi database Anda.

4. Jalankan aplikasi:

```bash
go run main.go
```

Server akan berjalan di `http://localhost:3000` (atau sesuai PORT yang dikonfigurasi).

---

## Environment Variables

```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=5432
JWT_SECRET_KEY=your_jwt_secret_key
JWT_EXPIRES_IN=24h
PORT=3000
```

---

## API Endpoints

### Autentikasi

- `POST /api/auth/register` - Registrasi user baru
- `POST /api/auth/login` - Login user
- `POST /api/auth/logout` - Logout user

### User

- `GET /api/users/me` - Mendapatkan profil user (Protected)

### Notes

- `POST /api/notes` - Membuat catatan baru (Protected)
- `GET /api/notes` - Mendapatkan semua catatan user (Protected)
- `GET /api/notes/:id` - Mendapatkan detail catatan (Protected)
- `PUT /api/notes/:id` - Update catatan (Protected)
- `DELETE /api/notes/:id` - Hapus catatan (Protected)

---

## Dokumentasi API

Dokumentasi lengkap API tersedia di folder `docs/api/`:

- [Auth API](docs/api/auth.md)
- [Users API](docs/api/users.md)
- [Notes API](docs/api/notes.md)

---

## Struktur Project

```
├── config/          # Konfigurasi database dan environment
├── docs/            # Dokumentasi API
├── handlers/        # Request Handlers
├── middlewares/     # Middleware (auth)
├── models/          # Model basis data
├── routes/          # Definisi rute
├── utils/           # Util func (JWT, dll)
└── main.go          # Entry point aplikasi
```

### Arsitektur

Project ini menggunakan **Layered Architecture Pattern** untuk memisahkan tanggung jawab setiap komponen:

### Keuntungan

✅ **Maintainable**: Code terorganisir, mudah debug dan update  
✅ **Scalable**: Mudah tambah fitur tanpa merusak existing code  
✅ **Testable**: Setiap layer dapat di-test secara terpisah  
✅ **Reusable**: Component dapat digunakan di berbagai tempat

---

## Contoh Penggunaan

### Register User

```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Membuat Catatan (dengan JWT token)

```bash
curl -X POST http://localhost:3000/api/notes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token" \
  -d '{
    "title": "Catatan Pertama",
    "content": "Ini adalah isi catatan pertama saya."
  }'
```

---

## Development

Untuk development, Anda bisa menggunakan tools seperti:

- **Air** untuk hot reload: `go install github.com/cosmtrek/air@latest`
- **Postman** atau **Insomnia** untuk testing API

---
