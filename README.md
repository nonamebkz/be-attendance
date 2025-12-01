# Backend Attendance System

Sistem backend untuk manajemen absensi karyawan menggunakan Go, Fiber, dan PostgreSQL.

## 📋 Prerequisites

Sebelum memulai, pastikan Anda telah menginstall:

- **Go 1.25.0** atau lebih tinggi dan **PostgreSQL**.

### Tools yang Diperlukan

- **Go** (1.25.0+)
- **PostgreSQL** (12+)
- **golang-migrate** (untuk database migration)
- **Make** (opsional, untuk menggunakan Makefile)

### Install golang-migrate

```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/migrate

# Windows (PowerShell)
# Download dari https://github.com/golang-migrate/migrate/releases
# Extract dan tambahkan ke PATH, atau gunakan:
choco install golang-migrate

# Atau menggunakan go install (semua platform)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Install Make di Windows

Windows tidak memiliki Make secara default. Berikut beberapa cara untuk menginstall:

#### Opsi 1: Menggunakan Chocolatey (Recommended)

1. Install Chocolatey terlebih dahulu (jika belum):
   ```powershell
   # Buka PowerShell sebagai Administrator
   Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
   ```

2. Install Make:
   ```powershell
   choco install make
   ```

#### Opsi 2: Menggunakan GnuWin32

1. Download Make dari: http://gnuwin32.sourceforge.net/packages/make.htm
2. Install dan tambahkan ke PATH environment variable
3. Restart terminal/PowerShell

#### Opsi 3: Menggunakan WSL (Windows Subsystem for Linux)

1. Install WSL:
   ```powershell
   wsl --install
   ```
2. Setelah WSL terinstall, gunakan terminal WSL untuk menjalankan command Make

#### Opsi 4: Menggunakan Git Bash

Jika Anda sudah install Git for Windows, Git Bash sudah include Make. Gunakan Git Bash untuk menjalankan command Make.

**Catatan:** Jika tidak ingin install Make, Anda bisa menggunakan command langsung tanpa Make (lihat bagian [Alternatif Tanpa Make](#alternatif-tanpa-make-windows)).

## 🚀 First Setup

### 1. Clone Repository

**Linux/macOS:**
```bash
git clone <repository-url>
cd be-attendance
```

**Windows (PowerShell/CMD):**
```cmd
git clone <repository-url>
cd be-attendance
```

**Windows (Git Bash):**
```bash
git clone <repository-url>
cd be-attendance
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Environment Variables

**Linux/macOS:**
```bash
cp example.env .env
```

**Windows (PowerShell):**
```powershell
Copy-Item example.env .env
```

**Windows (CMD):**
```cmd
copy example.env .env
```

**Windows (Git Bash):**
```bash
cp example.env .env
```

Edit file `.env` dengan konfigurasi database Anda:

```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=absensi_db
```

### 4. Setup Database

**Linux/macOS:**
```bash
# Masuk ke PostgreSQL
psql -U postgres

# Buat database
CREATE DATABASE absensi_db;
\q
```

**Windows:**
```cmd
# Masuk ke PostgreSQL (gunakan psql yang ada di PostgreSQL installation)
psql -U postgres

# Buat database
CREATE DATABASE absensi_db;
\q
```

Atau gunakan pgAdmin untuk membuat database secara visual.

### 5. Run Migrations

**Dengan Make:**
```bash
make migrate_up
```

**Tanpa Make (Windows - langsung menggunakan migrate):**
```cmd
# Pastikan migrate sudah di PATH, atau gunakan full path
migrate -path database/migration -database "postgresql://postgres:your_password@localhost:5432/absensi_db?sslmode=disable" -verbose up
```

### 6. Run Application

```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:3000`

## 📝 Makefile Commands

Makefile menyediakan beberapa command untuk memudahkan development:

> **Catatan untuk Windows:** Jika Make tidak terinstall, lihat bagian [Alternatif Tanpa Make](#alternatif-tanpa-make-windows) di bawah.

### Migration Commands

#### Create Migration

Membuat file migration baru:

```bash
make create_migration name=<nama_migration>
```

**Contoh:**
```bash
# Membuat tabel baru
make create_migration name=create_departemen_table

# Alter tabel yang sudah ada
make create_migration name=alter_karyawan_add_email
```

**Konvensi Naming:**
- **Create tabel baru**: Gunakan prefix `create_` diikuti nama tabel
  - Contoh: `create_departemen_table`, `create_lokasi_table`
- **Alter tabel**: Gunakan prefix `alter_` diikuti nama tabel dan perubahan
  - Contoh: `alter_karyawan_add_email`, `alter_absensi_add_keterangan`

#### Run Migrations Up

Menjalankan semua migration yang belum dijalankan:

```bash
make migrate_up
```

#### Run Migrations Down

Mengembalikan migration terakhir (rollback):

```bash
make migrate_down
```

#### Force Migration Version

Memaksa migration ke versi tertentu (hati-hati dengan command ini):

```bash
make force_migration version=<version_number>
```

**Contoh:**
```bash
make force_migration version=5
```

### Swagger Documentation

Generate Swagger documentation:

```bash
make swagger
```

### Clean

Menghapus file yang di-generate (seperti docs):

```bash
make clean
```

## 📚 Contoh Penggunaan

### Membuat Migration Baru

#### Contoh 1: Membuat Tabel Baru

```bash
make create_migration name=create_departemen_table
```

File yang dibuat:
- `database/migration/000008_create_departemen_table.up.sql`
- `database/migration/000008_create_departemen_table.down.sql`

Isi file `.up.sql`:
```sql
CREATE TABLE departemen (
    id_departemen SERIAL PRIMARY KEY,
    nama_departemen VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

Isi file `.down.sql`:
```sql
DROP TABLE IF EXISTS departemen;
```

#### Contoh 2: Alter Tabel yang Sudah Ada

```bash
make create_migration name=alter_karyawan_add_email
```

File yang dibuat:
- `database/migration/000009_alter_karyawan_add_email.up.sql`
- `database/migration/000009_alter_karyawan_add_email.down.sql`

Isi file `.up.sql`:
```sql
ALTER TABLE karyawan 
ADD COLUMN email VARCHAR(100) UNIQUE;
```

Isi file `.down.sql`:
```sql
ALTER TABLE karyawan 
DROP COLUMN IF EXISTS email;
```

### Menjalankan Migration

Setelah membuat migration, jalankan:

```bash
make migrate_up
```

### Rollback Migration

Jika perlu mengembalikan migration terakhir:

```bash
make migrate_down
```

## 🏗️ Project Structure

```
be-attendance/
├── config/              # Konfigurasi aplikasi
│   └── config.go
├── database/
│   ├── migration/      # Database migration files
│   └── uow/           # Unit of Work pattern
├── models/            # Data models
├── repositories/      # Data access layer
├── routes/            # API routes
├── main.go            # Entry point
├── Makefile           # Build automation
├── go.mod             # Go dependencies
└── README.md          # Dokumentasi
```

## 🔧 Development

### Menjalankan Development Server

```bash
go run main.go
```

### Build Binary

```bash
go build -o bin/app main.go
```

### Run Tests

```bash
go test ./...
```

## 🪟 Alternatif Tanpa Make (Windows)

Jika Anda tidak ingin menginstall Make di Windows, Anda bisa menjalankan command langsung:

### Create Migration

```cmd
migrate create -ext sql -dir database/migration -seq create_departemen_table
```

### Run Migrations Up

```cmd
migrate -path database/migration -database "postgresql://%DB_USER%:%DB_PASSWORD%@%DB_HOST%:%DB_PORT%/%DB_NAME%?sslmode=disable" -verbose up
```

Atau dengan nilai langsung:
```cmd
migrate -path database/migration -database "postgresql://postgres:password@localhost:5432/absensi_db?sslmode=disable" -verbose up
```

### Run Migrations Down

```cmd
migrate -path database/migration -database "postgresql://%DB_USER%:%DB_PASSWORD%@%DB_HOST%:%DB_PORT%/%DB_NAME%?sslmode=disable" -verbose down
```

### Force Migration

```cmd
migrate -path database/migration -database "postgresql://%DB_USER%:%DB_PASSWORD%@%DB_HOST%:%DB_PORT%/%DB_NAME%?sslmode=disable" force 5
```

### Membuat Batch File (Windows)

Anda bisa membuat file `migrate-up.bat` untuk memudahkan:

```batch
@echo off
setlocal
if exist .env (
    for /f "tokens=1,2 delims==" %%a in (.env) do set %%a=%%b
)
migrate -path database/migration -database "postgresql://%DB_USER%:%DB_PASSWORD%@%DB_HOST%:%DB_PORT%/%DB_NAME%?sslmode=disable" -verbose up
```

Kemudian jalankan:
```cmd
migrate-up.bat
```

## 📖 Teknologi yang Digunakan

- **Go 1.25.0** - Bahasa pemrograman
- **Fiber v2** - Web framework
- **sqlx** - SQL toolkit untuk Go
- **PostgreSQL** - Database
- **golang-migrate** - Database migration tool

## 🤝 Contributing

1. Buat branch baru untuk fitur (`git checkout -b feature/AmazingFeature`)
2. Commit perubahan Anda (`git commit -m 'Add some AmazingFeature'`)
3. Push ke branch (`git push origin feature/AmazingFeature`)
4. Buat Pull Request

## 🪟 Windows-Specific Notes

### Menjalankan Command di Windows

#### PowerShell
PowerShell adalah terminal modern yang direkomendasikan untuk Windows. Semua command Go dan migrate bisa dijalankan di PowerShell.

#### CMD (Command Prompt)
Command Prompt klasik juga bisa digunakan, tapi beberapa fitur mungkin terbatas.

#### Git Bash
Git Bash menyediakan environment seperti Linux/macOS dan sudah include Make jika Anda install Git for Windows.

### Troubleshooting Windows

**Problem: `make: command not found`**
- **Solusi:** Install Make menggunakan salah satu metode di atas, atau gunakan alternatif tanpa Make

**Problem: `migrate: command not found`**
- **Solusi:** Pastikan migrate sudah di PATH, atau gunakan full path ke executable migrate

**Problem: Environment variables tidak terbaca**
- **Solusi:** Pastikan file `.env` ada di root project dan formatnya benar (tanpa spasi di sekitar `=`)

**Problem: PostgreSQL connection refused**
- **Solusi:** 
  1. Pastikan PostgreSQL service berjalan
  2. Cek firewall Windows
  3. Pastikan `DB_HOST` di `.env` adalah `localhost` atau `127.0.0.1`

### Recommended Tools untuk Windows Development

- **VS Code** - Editor code dengan Go extension
- **Git for Windows** - Include Git Bash yang sudah ada Make
- **PostgreSQL for Windows** - Database server
- **DBeaver** atau **pgAdmin** - Database management tools

## 📄 License

[Your License Here]

